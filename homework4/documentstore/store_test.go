package documentstore

import (
	"fmt"
	"testing"
)

func TestNewStore(t *testing.T) {
	t.Run("Should create empty store", func(t *testing.T) {
		store := NewStore()
		if len(store.collections) != 0 {
			t.Error(fmt.Errorf("should create store with no collections: %v", store.collections))
		}
	})
}

func TestStore_CreateCollection(t *testing.T) {
	t.Run("Should add collection to store", func(t *testing.T) {
		store := NewStore()
		if ok, collection := store.CreateCollection("test", &CollectionConfig{PrimaryKey: "primaryKey"}); !ok {
			t.Error(fmt.Errorf("should return 'true' on new collection creation, got: %d", len(store.collections)))
		} else if collection == nil {
			t.Error(fmt.Errorf("returned nil collection"))
		}

		if len(store.collections) != 1 {
			t.Error(fmt.Errorf("store should have 1 collection, got: %d", len(store.collections)))
		}
	})

	t.Run("Should not add duplicated collection to store", func(t *testing.T) {
		store := NewStore()
		store.CreateCollection("test", &CollectionConfig{PrimaryKey: "primaryKey"})

		if created, _ := store.CreateCollection("test", &CollectionConfig{PrimaryKey: "primaryKey2"}); created {
			t.Error(fmt.Errorf("returned 'true' on duplicated collection createion: %v", store.collections))
		}

		if len(store.collections) != 1 {
			t.Error(fmt.Errorf("store should have 1 collection, got: %d", len(store.collections)))
		}
		if keyValue := store.collections["test"].PrimaryKey; keyValue != "primaryKey" {
			t.Error(fmt.Errorf("wrong collection is saved in store. Passed: %s", keyValue))
		}
	})
}

func TestStore_GetCollection(t *testing.T) {
	t.Run("Should get document from store", func(t *testing.T) {
		store := NewStore()
		store.CreateCollection("test", &CollectionConfig{PrimaryKey: "primaryKey"})
		collection, ok := store.GetCollection("test")
		if !ok {
			t.Error(fmt.Errorf("store should have a collection, got: %v", collection))
		}
		if collection.CollectionConfig.PrimaryKey != "primaryKey" {
			t.Error(fmt.Errorf("wrong collection is saved in store. Passed: %v", collection))
		}
	})
}

func TestStore_DeleteCollection(t *testing.T) {
	t.Run("Should delete document from store", func(t *testing.T) {
		store := NewStore()
		store.CreateCollection("test", &CollectionConfig{PrimaryKey: "primaryKey"})
		if ok := store.DeleteCollection("test"); !ok {
			t.Error(fmt.Errorf("returned 'false' on collection deletion"))
		}
		if collection, ok := store.GetCollection("test"); ok {
			t.Error(fmt.Errorf("collection was not deleted from the store: %v", collection))
		}
	})

	t.Run("Should return false if collection was not deleted", func(t *testing.T) {
		store := NewStore()
		if ok := store.DeleteCollection("test"); ok {
			t.Error(fmt.Errorf("returned 'true' in invalid collection deletion: %v", store))
		}
	})
}

func TestStore_DocumentCreation(t *testing.T) {
	t.Run("Should add document in store", func(t *testing.T) {
		store := NewStore()
		_, collection := store.CreateCollection("test", &CollectionConfig{PrimaryKey: "primaryKey"})
		primaryKey := "uniqueKye"
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"primaryKey": {Type: DocumentFieldTypeString, Value: primaryKey},
			},
		})
		if col, ok := collection.Get(primaryKey); !ok {
			t.Error(fmt.Errorf("collection should have a document with primaryKey: %v", col))
		}
	})
}
