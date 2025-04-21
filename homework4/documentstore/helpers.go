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

func ValidateDocument(documentKey string, doc Document) (DocumentValidatorErrors, bool) {
	validatorErrors := DocumentValidatorErrors{}

	if documentConfigValue, ok := doc.Fields[documentKey]; !ok {
		validatorErrors = append(validatorErrors, fmt.Errorf("config key is missing: %s", documentKey))
	} else if GetType(documentConfigValue.Value) != DocumentFieldTypeString {
		validatorErrors = append(
			validatorErrors,
			fmt.Errorf("config key type is incorrect. Should be string, passed: %s", GetType(documentConfigValue)),
		)
	}

	for key, value := range doc.Fields {
		var errorMsg string
		if t := GetType(value.Value); t != value.Type {
			errorMsg = fmt.Sprintf("Document field %s type mismatch. Expected: %s, got: %s", key, value.Type, t)
			validatorErrors = append(validatorErrors, errors.New(errorMsg))
		}

	}

	return validatorErrors, len(validatorErrors) == 0
}
