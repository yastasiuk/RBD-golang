package documentstore

import (
	"errors"
	"fmt"
	"log/slog"
	"reflect"
)

var ErrDocumentNotFound = errors.New("document not found")
var ErrValidationFailed = errors.New("validation failed")

var ErrUserUniquenessValidation = errors.New("user already exists")

type Collection struct {
	Cfg       CollectionConfig    `json:"cfg"`
	Documents map[string]Document `json:"documents"`
}

type CollectionConfig struct {
	PrimaryKey string `json:"primaryKey"`
}

func NewCollection(cfg *CollectionConfig) *Collection {
	col := Collection{
		Cfg:       *cfg,
		Documents: map[string]Document{},
	}

	return &col
}

func (s *Collection) Put(doc Document) (*Document, error) {
	slog.Debug("Put document", "doc", doc)
	primaryKeyField, ok := doc.Fields[s.Cfg.PrimaryKey]
	if !ok {
		return nil, fmt.Errorf("%w: PrimaryKey is missing Fields", ErrValidationFailed)
	}

	id, typedCorrectly := primaryKeyField.Value.(string)

	if !typedCorrectly {
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

	s.Documents[id] = doc
	return &doc, nil
}

func (s *Collection) Get(key string) (*Document, error) {
	slog.Debug("Get document:", "key", key)
	if doc, ok := s.Documents[key]; ok {
		return &doc, nil
	}

	return nil, fmt.Errorf("failed to find document with key %s: %w", key, ErrDocumentNotFound)
}

func (s *Collection) Delete(key string) bool {
	slog.Debug("Delete document:", "key", key)
	if _, err := s.Get(key); err == nil {
		delete(s.Documents, key)
		return true
	}

	return false
}

func (s *Collection) List() []Document {
	slog.Debug("List documents in collection")
	docs := make([]Document, 0, len(s.Documents))
	for _, doc := range s.Documents {
		docs = append(docs, doc)
	}

	return docs
}
