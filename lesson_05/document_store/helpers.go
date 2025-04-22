package documentstore

import (
	"errors"
	"fmt"
)

type DocumentValidatorErrors = []error

func GetType(k interface{}) DocumentFieldType {
	switch k.(type) {
	case string:
		return DocumentFieldTypeString
	case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64, float32, float64, uintptr:
		return DocumentFieldTypeNumber
	case bool:
		return DocumentFieldTypeBool
	case []interface{}:
		return DocumentFieldTypeArray
	case map[string]interface{}:
		return DocumentFieldTypeObject
	default:
		return "unknown"
	}
}

func validateDocument(doc Document) error {
	validatorErrors := DocumentValidatorErrors{}
	for key, value := range doc.Fields {
		var errorMsg string
		if t := GetType(value.Value); t != value.Type {
			errorMsg = fmt.Sprintf("Document field %s type mismatch. Expected: %s, got: %s", key, value.Type, t)
			validatorErrors = append(validatorErrors, errors.New(errorMsg))
		}
	}

	if len(validatorErrors) > 0 {
		return errors.Join(validatorErrors...)
	}

	return nil
}
