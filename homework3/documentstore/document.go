package documentstore

import (
	"errors"
	"fmt"
	"reflect"
	"slices"
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

var numberTypes = []reflect.Kind{
	reflect.Int,
	reflect.Int8,
	reflect.Int16,
	reflect.Int32,
	reflect.Int64,
	reflect.Uint,
	reflect.Uint8,
	reflect.Uint16,
	reflect.Uint32,
	reflect.Uint64,
	reflect.Uintptr,
	reflect.Float32,
	reflect.Float64,
	reflect.Complex64,
	reflect.Complex128,
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
		switch value.Type {
		case DocumentFieldTypeString:
			if t := reflect.TypeOf(value.Value).Kind(); t != reflect.String {
				errorMsg = fmt.Sprintf("Document field %s type mismatch: expected string, got %s", key, t)
			}
		case DocumentFieldTypeNumber:
			if t := reflect.TypeOf(value.Value).Kind(); !slices.Contains(numberTypes, t) {
				errorMsg = fmt.Sprintf("Document field %s type mismatch: expected number, got %s", key, t)
			}
		case DocumentFieldTypeBool:
			if t := reflect.TypeOf(value.Value).Kind(); t != reflect.Bool {
				errorMsg = fmt.Sprintf("Document field %s type mismatch: expected boolean, got %s", key, t)
			}
		case DocumentFieldTypeArray:
			if t := reflect.TypeOf(value.Value).Kind(); t != reflect.Array && t != reflect.Slice {
				errorMsg = fmt.Sprintf("Document field %s type mismatch: expected array, got %s", key, t)
			}
		case DocumentFieldTypeObject:
			if t := reflect.TypeOf(value.Value).Kind(); t != reflect.Struct {
				errorMsg = fmt.Sprintf("Document field %s type mismatch: expected object, got %s", key, t)
			}
		default:
			errorMsg = fmt.Sprintf("Document field %s has unsupported type %s", key, value.Type)
		}

		if errorMsg != "" {
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
