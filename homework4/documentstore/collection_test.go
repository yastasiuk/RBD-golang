package documentstore

import (
	"fmt"
	"testing"
)

func TestCollection_CreateNew(t *testing.T) {
	t.Run("Should create new collection", func(t *testing.T) {
		primaryKey := "primaryKey"
		if collection := NewCollection(&CollectionConfig{PrimaryKey: primaryKey}); collection.PrimaryKey != primaryKey {
			t.Error(fmt.Errorf("expected collection to be created with primaryKey %s, passed %s", primaryKey, primaryKey))
		}
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

		if len(collection.documents) != 1 {
			t.Error(fmt.Errorf("store should return 1 elements. Returned: %d", len(collection.documents)))
		}
	})

	t.Run("Should not add invalid document", func(t *testing.T) {
		collection := NewCollection(&CollectionConfig{PrimaryKey: "primaryKey"})
		collection.Put(Document{})
		if len(collection.documents) != 0 {
			t.Error(fmt.Errorf("store should return 0 elements. Returned: %d", len(collection.documents)))
		}
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
		if len(collection.documents) != 1 {
			t.Error(fmt.Errorf("store should return 1 elements. Returned: %d", len(collection.documents)))
		}
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
		if ok := collection.Delete("123"); !ok {
			t.Error(fmt.Errorf("collection should return true after deleting an existing document"))
		}
	})

	t.Run("Should return false if document was not deleted", func(t *testing.T) {
		collection := NewCollection(&CollectionConfig{PrimaryKey: "primaryKey"})
		if ok := collection.Delete("123"); ok {
			t.Error(fmt.Errorf("collection should return false if document was not deleted"))
		}
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
		if len(docs) != 1 {
			t.Error(fmt.Errorf("store should return 1 elements. Returned: %d", len(docs)))
		}
	})
}
