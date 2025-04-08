package documentstore

import (
	"errors"
	"fmt"
)

type DocumentValidatorErrors = []error

var requiredFields = []string{
	"key",
}

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

func ValidateDocument(doc Document) (DocumentValidatorErrors, bool) {
	requiredFieldsCheck := map[string]interface{}{}
	for _, fieldKey := range requiredFields {
		requiredFieldsCheck[fieldKey] = nil
	}

	validatorErrors := DocumentValidatorErrors{}

	for key, value := range doc.Fields {
		if _, ok := requiredFieldsCheck[key]; ok {
			delete(requiredFieldsCheck, key)
		}

		var errorMsg string
		if t := GetType(value.Value); t != value.Type {
			errorMsg = fmt.Sprintf("Document field %s type mismatch. Expected: %s, got: %s", key, value.Type, t)
			validatorErrors = append(validatorErrors, errors.New(errorMsg))
		}

	}

	for _, missingRequiredField := range requiredFieldsCheck {
		validatorErrors = append(validatorErrors, errors.New(fmt.Sprintf("field %d is required", missingRequiredField)))
	}

	return validatorErrors, len(validatorErrors) == 0
}

func GetDocumentKey(doc Document) (string, bool) {
	for key, value := range doc.Fields {
		if key == "key" {
			return value.Value.(string), true
		}
	}

	return "", false
}
