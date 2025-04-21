package documentstore

import "fmt"

type Collection struct {
	CollectionConfig
	documents []Document
}

type CollectionConfig struct {
	PrimaryKey string
}

func NewCollection(cfg *CollectionConfig) *Collection {
	col := Collection{
		CollectionConfig: *cfg,
		documents:        []Document{},
	}

	return &col
}

func (s *Collection) Put(doc Document) {
	// Потрібно перевірити що документ містить поле `{cfg.PrimaryKey}` типу `string`
	if err, ok := ValidateDocument(s.PrimaryKey, doc); !ok {
		fmt.Println("Document is not valid:", err)
		return
	}

	if _, exists := s.Get(doc.Fields[s.PrimaryKey].Value.(string)); exists {
		fmt.Println("Document already exists:", doc.Fields[s.PrimaryKey].Value)
		return
	}

	s.documents = append(s.documents, doc)
}

func (s *Collection) Get(key string) (*Document, bool) {
	for _, doc := range s.documents {
		if doc.Fields[s.PrimaryKey].Value == key {
			return &doc, true
		}
	}

	return nil, false
}

func (s *Collection) Delete(key string) bool {
	for i, doc := range s.documents {
		if doc.Fields[s.PrimaryKey].Value == key {
			sls := make([]Document, len(s.documents)-1)
			copy(sls, append(s.documents[:i], s.documents[i+1:]...))
			s.documents = sls
			return true
		}
	}

	return false
}

func (s *Collection) List() []Document {
	return s.documents
}
