package validator

import (
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	"regexp"
)

func registerSlugValidation(val *validator.Validate, trans ut.Translator) error {
	if err := val.RegisterValidation("slug",
		func(fl validator.FieldLevel) bool {
			slug := fl.Field().String()
			// Regular expression to match lowercase alphanumeric characters and hyphens
			re := regexp.MustCompile(`^[a-z0-9-]+$`)
			return re.MatchString(slug)
		},
	); err != nil {
		return err
	}

	if err := val.RegisterTranslation("slug", trans,
		func(ut ut.Translator) error {
			return ut.Add("slug", "{0} harus berupa slug (hanya berisi kata-kata non-kapital, angka, dan tanda hubung)", true)
		},

		func(ut ut.Translator, fe validator.FieldError) string {
			t, _ := ut.T("slug", fe.Field())
			return t
		},
	); err != nil {
		return err
	}

	return nil
}
