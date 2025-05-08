package documentstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCollection_CreateNew(t *testing.T) {
	t.Run("Should create new collection", func(t *testing.T) {
		primaryKey := "primaryKey"
		collection := NewCollection(&CollectionConfig{PrimaryKey: primaryKey})
		assert.Equal(t, primaryKey, collection.Cfg.PrimaryKey)
	})
}

func TestCollection_Put(t *testing.T) {
	t.Run("Should added document to a collection", func(t *testing.T) {
		collection := NewCollection(&CollectionConfig{PrimaryKey: "primaryKey"})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"primaryKey": {Value: "123", Type: DocumentFieldTypeString},
			},
		})

		assert.Equal(t, 1, len(collection.Documents))
	})

	t.Run("Should not add invalid document", func(t *testing.T) {
		collection := NewCollection(&CollectionConfig{PrimaryKey: "primaryKey"})
		collection.Put(Document{})
		assert.Equal(t, 0, len(collection.Documents))
	})

	t.Run("Should not add already existing document", func(t *testing.T) {
		collection := NewCollection(&CollectionConfig{PrimaryKey: "primaryKey"})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"primaryKey": {Type: DocumentFieldTypeString, Value: "123"},
			},
		})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"primaryKey": {Type: DocumentFieldTypeString, Value: "123"},
			},
		})
		assert.Equal(t, 1, len(collection.Documents))
	})
}

func TestCollection_Delete(t *testing.T) {
	t.Run("Should delete document in collection", func(t *testing.T) {
		collection := NewCollection(&CollectionConfig{PrimaryKey: "primaryKey"})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"primaryKey": {Type: DocumentFieldTypeString, Value: "123"},
			},
		})
		ok := collection.Delete("123")
		assert.Equal(t, true, ok, "Should delete document in collection")
	})

	t.Run("Should return false if document was not deleted", func(t *testing.T) {
		collection := NewCollection(&CollectionConfig{PrimaryKey: "primaryKey"})
		ok := collection.Delete("123")
		assert.Equal(t, false, ok, "collection should return false if document was not deleted")
	})
}

func TestCollection_List(t *testing.T) {
	t.Run("Should list documents in collection", func(t *testing.T) {
		collection := NewCollection(&CollectionConfig{PrimaryKey: "primaryKey"})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"primaryKey": {Type: DocumentFieldTypeString, Value: "123"},
			},
		})
		docs := collection.List()
		assert.Equal(t, 1, len(docs), "collection should return false if document was not deleted")
	})
}
