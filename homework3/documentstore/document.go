package documentstore

import (
	"errors"
	"fmt"
)

type DocumentFieldType string

const (
	DocumentFieldTypeString DocumentFieldType = "string"
	DocumentFieldTypeNumber DocumentFieldType = "number"
	DocumentFieldTypeBool   DocumentFieldType = "bool"
	DocumentFieldTypeArray  DocumentFieldType = "array"
	DocumentFieldTypeObject DocumentFieldType = "object"
)

type DocumentField struct {
	Type  DocumentFieldType
	Value interface{}
}

type Document struct {
	Fields map[string]DocumentField
}

type DocumentValidatorErrors = []error

var requiredFields = []string{
	"key",
}

func getType(k interface{}) DocumentFieldType {
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

func validateDocument(doc Document) (DocumentValidatorErrors, bool) {
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
		if t := getType(value.Value); t != value.Type {
			errorMsg = fmt.Sprintf("Document field %s type mismatch. Expected: %s, got: %s", key, value.Type, t)
			validatorErrors = append(validatorErrors, errors.New(errorMsg))
		}

	}

	for _, missingRequiredField := range requiredFieldsCheck {
		validatorErrors = append(validatorErrors, errors.New(fmt.Sprintf("field %d is required", missingRequiredField)))
	}

	return validatorErrors, len(validatorErrors) == 0
}

func getDocumentKey(doc Document) (string, bool) {
	for key, value := range doc.Fields {
		if key == "key" {
			return value.Value.(string), true
		}
	}

	return "", false
}

var documents = map[string]Document{}

func Put(doc Document) {
	// 1. Перевірити що документ містить в мапі поле `key` типу `string`
	// 2. Додати Document до локальної мапи з документами
	err, valid := validateDocument(doc)
	if !valid {
		fmt.Println("Document is not valid:", err)
		return
	}

	if key, ok := getDocumentKey(doc); !ok {
		fmt.Println("Key value is missing for doc:", doc)
	} else {
		documents[key] = doc
	}
}

func Get(key string) (*Document, bool) {
	// Потрібно повернути документ по ключу
	// Якщо документ знайдено, повертаємо `true` та поінтер на документ
	// Інакше повертаємо `false` та `nil`
	if doc, ok := documents[key]; !ok {
		return nil, false
	} else {
		return &doc, ok
	}
}

func Delete(key string) bool {
	// Видаляємо документа по ключу.
	// Повертаємо `true` якщо ми знайшли і видалили документі
	// Повертаємо `false` якщо документ не знайдено
	_, ok := documents[key]
	if !ok {
		return false
	} else {
		delete(documents, key)
		return true
	}
}

func List() []Document {
	// Повертаємо список усіх документів
	docs := make([]Document, 0, len(documents))

	for _, doc := range documents {
		docs = append(docs, doc)
	}

	return docs
}

func ResetStore() {
	documents = map[string]Document{}
}
