package documentstore

import (
	"errors"
	"fmt"
	"reflect"
)

type DocumentValidatorErrors = []error

func GetType(k reflect.Kind) DocumentFieldType {
	switch k {
	case reflect.String:
		return DocumentFieldTypeString
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64, reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Float32, reflect.Float64, reflect.Uintptr, reflect.Complex64, reflect.Complex128:
		return DocumentFieldTypeNumber
	case reflect.Bool:
		return DocumentFieldTypeBool
	case reflect.Array, reflect.Slice:
		return DocumentFieldTypeArray
	case reflect.Map, reflect.Struct:
		return DocumentFieldTypeObject
	default:
		return "unknown"
	}
}

func validateDocument(doc Document) error {
	validatorErrors := DocumentValidatorErrors{}
	for key, value := range doc.Fields {
		var errorMsg string
		if t := GetType(reflect.TypeOf(value.Value).Kind()); t != value.Type {
			errorMsg = fmt.Sprintf("Document field %s type mismatch. Expected: %s, got: %s", key, value.Type, t)
			validatorErrors = append(validatorErrors, errors.New(errorMsg))
		}
	}

	if len(validatorErrors) > 0 {
		return errors.Join(validatorErrors...)
	}

	return nil
}
