package validator

import (
	"errors"
	"fmt"
	"github.com/bem-filkom/web-bem-backend/internal/pkg/log"
	"github.com/go-playground/validator/v10"
	jsoniter "github.com/json-iterator/go"
	"reflect"
	"sync"

	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	idtranslations "github.com/go-playground/validator/v10/translations/id"
)

var validatorObj *ValidatorStruct

type ValidatorStruct struct {
	validator  *validator.Validate
	translator ut.Translator
}

type validationError struct {
	Tag         string `json:"tag"`
	Param       string `json:"param"`
	Translation string `json:"translation"`
}

type ValidationErrorsResponse []map[string]validationError

var once sync.Once

func (v ValidationErrorsResponse) Error() string {
	j, err := jsoniter.Marshal(v)
	if err != nil {
		return ""
	}

	return string(j)
}

func (v ValidationErrorsResponse) Serialize() any {
	return v
}

func createValidator() {
	fmt.Println("Creating validator instance")

	idInstance := id.New()
	uni := ut.New(idInstance, idInstance)

	translator, _ := uni.GetTranslator("id")

	val := validator.New()
	err := idtranslations.RegisterDefaultTranslations(val, translator)
	if err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error": err.Error(),
		}).Fatalln("[VALIDATOR][NewValidator] Failed to register default translations")
		return
	}

	if err := registerCustomValidations(val, translator); err != nil {
		log.GetLogger().WithFields(map[string]any{
			"error": err.Error(),
		}).Fatalln("[VALIDATOR][NewValidator] Failed to register custom validations")
		return
	}

	validatorObj = &ValidatorStruct{
		validator:  val,
		translator: translator,
	}
}

func registerCustomValidations(val *validator.Validate, trans ut.Translator) error {
	if err := registerSlugValidation(val, trans); err != nil {
		return err
	}

	return nil
}

func GetValidator() *ValidatorStruct {
	once.Do(createValidator)

	return validatorObj
}

func getJSONFieldName(field reflect.StructField) string {
	checkTags := []string{"json", "query", "param"}
	for _, tag := range checkTags {
		jsonTag := field.Tag.Get(tag)
		if jsonTag != "" {
			return jsonTag
		}
	}

	return field.Name
}

func (v *ValidatorStruct) ValidateStruct(data interface{}) ValidationErrorsResponse {
	if err := v.validator.Struct(data); err != nil {
		return v.handleValidationErrors(err, data)
	}
	return nil
}

func (v *ValidatorStruct) ValidateVariable(data interface{}, tag string) ValidationErrorsResponse {
	if err := v.validator.Var(data, tag); err != nil {
		return v.handleValidationErrors(err, nil)
	}
	return nil
}

func (v *ValidatorStruct) handleValidationErrors(err error, data interface{}) ValidationErrorsResponse {
	var valErrs validator.ValidationErrors
	if errors.As(err, &valErrs) {
		length := len(valErrs)
		res := make(ValidationErrorsResponse, length)
		for i, err := range valErrs {
			jsonTag := ""
			if data != nil {
				dataType := reflect.TypeOf(data)
				if dataType.Kind() == reflect.Ptr {
					dataType = dataType.Elem()
				}
				field, _ := dataType.FieldByName(err.StructField())
				jsonTag = getJSONFieldName(field)
			}

			res[i] = map[string]validationError{
				jsonTag: {
					Tag:         err.Tag(),
					Param:       err.Param(),
					Translation: err.Translate(v.translator),
				},
			}
		}
		return res
	}

	log.GetLogger().WithFields(map[string]any{
		"error": err.Error(),
	}).Error("[VALIDATOR] Validation failed")
	return nil
}
