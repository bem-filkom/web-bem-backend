package sqlutil

import (
	"fmt"
	"github.com/google/uuid"
	"reflect"
	"strings"
)

type SQLUpdateBuilder struct {
	QueryParts []string
	Args       []any
	ArgIndex   int
}

// newSQLUpdateBuilder creates a new SQLUpdateBuilder instance
func newSQLUpdateBuilder() *SQLUpdateBuilder {
	return &SQLUpdateBuilder{
		ArgIndex: 1,
	}
}

func GenerateUpdateQueryPart(dto any) (string, []any, int, error) {
	builder := newSQLUpdateBuilder()
	v := reflect.ValueOf(dto)
	if v.Kind() != reflect.Ptr && v.Kind() != reflect.Struct {
		return "", nil, -1, fmt.Errorf("dto must be a struct or pointer to struct")
	}

	if v.Kind() == reflect.Ptr {
		v = v.Elem()
		if v.Kind() != reflect.Struct {
			return "", nil, -1, fmt.Errorf("dto must be a pointer to struct")
		}
	}

	t := v.Type()
	for i := 0; i < v.NumField(); i++ {
		fieldValue := v.Field(i)
		fieldType := t.Field(i)
		fieldName := fieldType.Name

		// Get the database tag, fallback to lowercase field name, skip if tag is "-"
		dbTag := t.Field(i).Tag.Get("db")
		if dbTag == "-" {
			continue
		}
		if dbTag == "" {
			dbTag = strings.ToLower(fieldName)
		}

		switch fieldValue.Kind() {
		case reflect.Ptr:
			if fieldValue.IsNil() {
				continue
			}
		case reflect.Slice, reflect.Map, reflect.Interface, reflect.Chan:
			if fieldValue.IsNil() {
				continue
			}
		default:
			switch fieldType.Type {
			case reflect.TypeOf(uuid.UUID{}):
				if fieldValue.Interface() == uuid.Nil {
					continue
				}
			}
		}

		builder.addField(dbTag, fieldValue.Interface())
	}

	if len(builder.QueryParts) == 0 {
		return "", nil, -1, fmt.Errorf("no fields to update")
	}

	return strings.Join(builder.QueryParts, ", "), builder.Args, builder.ArgIndex, nil
}

// addField adds a field to the update query
func (b *SQLUpdateBuilder) addField(fieldName string, value interface{}) {
	b.QueryParts = append(b.QueryParts, fmt.Sprintf("%s = $%d", fieldName, b.ArgIndex))
	b.Args = append(b.Args, value)
	b.ArgIndex++
}
