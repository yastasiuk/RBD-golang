package documentstore

import "fmt"

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

func (s *Collection) Put(doc Document) {
	primaryKeyField, ok := doc.Fields[s.cfg.PrimaryKey]
	if !ok {
		fmt.Printf("Config key is missing in doc, %v\n", doc)
		return
	}

	id, typedCorrectly := primaryKeyField.Value.(string)

	if !typedCorrectly {
		fmt.Println("Config key has incorrect type")
		return
	}

	if err, ok := validateDocument(doc); !ok {
		fmt.Println("Document is not valid:", err)
		return
	}

	s.documents[id] = doc
}

func (s *Collection) Get(key string) (*Document, bool) {
	if doc, ok := s.documents[key]; ok {
		return &doc, true
	}

	return nil, false
}

func (s *Collection) Delete(key string) bool {
	if _, exists := s.Get(key); exists {
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
