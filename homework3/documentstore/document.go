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

func validateDocument(doc Document) (DocumentValidatorErrors, bool) {
	validatorErrors := DocumentValidatorErrors{}
	hasKeyField := false
	for key, value := range doc.Fields {
		if key == "key" {
			hasKeyField = true
			if value.Type != DocumentFieldTypeString {
				validatorErrors = append(validatorErrors, errors.New(fmt.Sprintf("field 'key' has invalid type. Required: 'string'. Passed: %s", value.Type)))
			}
			switch t := value.Value.(type) {
			case string:
				if value.Value.(string) == "" {
					validatorErrors = append(validatorErrors, errors.New("field 'key' value is required"))
				}
			default:
				validatorErrors = append(validatorErrors, errors.New(fmt.Sprintf("field 'key' has invalid value type. Required: 'string'. Passed: %s", t)))
			}
		}

	}

	if !hasKeyField {
		validatorErrors = append(validatorErrors, errors.New("field 'key' not found in document"))
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
