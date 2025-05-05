package documentstore

import (
	"errors"
	"fmt"
	"reflect"
)

var ErrDocumentNotFound = errors.New("document not found")
var ErrValidationFailed = errors.New("validation failed")

var ErrUserUniquenessValidation = errors.New("user already exists")

type Collection struct {
	cfg       CollectionConfig
	documents map[string]Document
}

type CollectionConfig struct {
	PrimaryKey string
}

func NewCollection(cfg *CollectionConfig) *Collection {
	col := Collection{
		cfg:       *cfg,
		documents: map[string]Document{},
	}

	return &col
}

func (s *Collection) Put(doc Document) (*Document, error) {
	primaryKeyField, ok := doc.Fields[s.cfg.PrimaryKey]
	if !ok {
		fmt.Printf("Config key is missing in doc, %v\n", doc)
		return nil, fmt.Errorf("%w: PrimaryKey is missing Fields", ErrValidationFailed)
	}

	id, typedCorrectly := primaryKeyField.Value.(string)

	if !typedCorrectly {
		fmt.Println("Config key has incorrect type", doc.Fields)
		return nil, fmt.Errorf("%w: PrimaryKey has incorrect type. Should be a string, passed %v", ErrValidationFailed, reflect.TypeOf(doc.Fields))
	}

	if id == "" {
		return nil, fmt.Errorf("%w: PrimaryKey cannot be empty", ErrValidationFailed)
	}

	if _, exists := s.Get(id); exists == nil {
		return nil, fmt.Errorf("%w; ID: '%s'", ErrUserUniquenessValidation, id)
	}

	if validationErrors := validateDocument(doc); validationErrors != nil {
		return nil, fmt.Errorf("%w: document validation failed: %w", ErrValidationFailed, validationErrors)
	}

	s.documents[id] = doc
	return &doc, nil
}

func (s *Collection) Get(key string) (*Document, error) {
	if doc, ok := s.documents[key]; ok {
		return &doc, nil
	}

	return nil, fmt.Errorf("failed to find document with key %s: %w", key, ErrDocumentNotFound)
}

func (s *Collection) Delete(key string) bool {
	if _, err := s.Get(key); err == nil {
		delete(s.documents, key)
		return true
	}

	return false
}

func (s *Collection) List() []Document {
	docs := make([]Document, 0, len(s.documents))
	for _, doc := range s.documents {
		docs = append(docs, doc)
	}

	return docs
}
