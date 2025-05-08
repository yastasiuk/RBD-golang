package documentstore

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestStore_NewStore(t *testing.T) {
	store := NewStore()
	assert.NotNil(t, store)
	assert.Empty(t, store.Collections)
}

func TestStore_CreateCollection(t *testing.T) {
	t.Run("Should create a new collection", func(t *testing.T) {
		store := NewStore()

		ok, col := store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
		assert.True(t, ok)
		assert.NotNil(t, col)

		retrieved, exists := store.GetCollection("users")
		assert.True(t, exists)
		assert.Equal(t, col, retrieved)
	})

	t.Run("Should not create duplicated collection", func(t *testing.T) {
		store := NewStore()
		store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})

		ok, col := store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
		assert.False(t, ok)
		assert.Nil(t, col)
	})
}

func TestStore_GetCollection(t *testing.T) {
	t.Run("Should return collection", func(t *testing.T) {
		store := NewStore()
		store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})

		_, exists := store.GetCollection("users")
		assert.True(t, exists)
	})
	t.Run("Should return not found error", func(t *testing.T) {
		store := NewStore()

		col, ok := store.GetCollection("notExistingCollection")
		assert.False(t, ok)
		assert.Nil(t, col)
	})
}

func TestDeleteCollection_Success(t *testing.T) {
	t.Run("Should delete collection", func(t *testing.T) {
		store := NewStore()
		store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})

		store.DeleteCollection("users")

		_, exists := store.GetCollection("users")
		assert.False(t, exists)
	})

	t.Run("Should return not found error", func(t *testing.T) {
		store := NewStore()

		ok := store.DeleteCollection("nonexistent")
		assert.False(t, ok)
	})
}

func TestStore_Dump(t *testing.T) {
	t.Run("Should dump", func(t *testing.T) {
		store := NewStore()
		_, collection := store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"id": {
					Type:  DocumentFieldTypeString,
					Value: "unique_id",
				},
				"name": {
					Type:  DocumentFieldTypeString,
					Value: "John Doe",
				},
			},
		})
		bytes, err := store.Dump()
		assert.Nil(t, err)
		expectedResult := `{
  "collections": {
    "users": {
      "cfg": {
        "primaryKey": "id"
      },
      "documents": {
        "unique_id": {
          "fields": {
            "id": {
              "type": "string",
              "value": "unique_id"
            },
            "name": {
              "type": "string",
              "value": "John Doe"
            }
          }
        }
      }
    }
  }
}`
		assert.Equal(t, expectedResult, string(bytes), "Should return expected string")
	})
}

func TestStore_NewStoreFromDump(t *testing.T) {
	t.Run("Should create a new store from dump", func(t *testing.T) {
		store := NewStore()
		_, collection := store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"id": {
					Type:  DocumentFieldTypeString,
					Value: "unique_id",
				},
				"name": {
					Type:  DocumentFieldTypeString,
					Value: "John Doe",
				},
			},
		})
		bytes, _ := store.Dump()
		store2, err := NewStoreFromDump(bytes)
		assert.Nil(t, err)
		assert.Equal(t, store, store2, "Stores should be the same")
	})

	t.Run("Should return error on invalid bytes", func(t *testing.T) {
		store, err := NewStoreFromDump([]byte("random bytes"))
		assert.NotNil(t, err)
		assert.Nil(t, store)
	})
}

func TestStore_DumpToFile(t *testing.T) {
	t.Run("Should dump to file", func(t *testing.T) {
		tmpFile := "store_test_dump.json"
		defer os.Remove(tmpFile)

		store := NewStore()
		_, collection := store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"id": {
					Type:  DocumentFieldTypeString,
					Value: "unique_id",
				},
				"name": {
					Type:  DocumentFieldTypeString,
					Value: "John Doe",
				},
			},
		})

		err := store.DumpToFile(tmpFile)
		assert.Nil(t, err)
		bytes, err := os.ReadFile(tmpFile)
		assert.Nil(t, err)
		expectedResult := `{
  "collections": {
    "users": {
      "cfg": {
        "primaryKey": "id"
      },
      "documents": {
        "unique_id": {
          "fields": {
            "id": {
              "type": "string",
              "value": "unique_id"
            },
            "name": {
              "type": "string",
              "value": "John Doe"
            }
          }
        }
      }
    }
  }
}`
		assert.Equal(t, expectedResult, string(bytes), "Should return expected string")
	})
}

func TestStore_NewStoreFromFile(t *testing.T) {
	t.Run("Should create a new store from file", func(t *testing.T) {
		tmpFile := "store_test_dump.json"
		defer os.Remove(tmpFile)
		store := NewStore()
		_, collection := store.CreateCollection("users", &CollectionConfig{PrimaryKey: "id"})
		collection.Put(Document{
			Fields: map[string]DocumentField{
				"id": {
					Type:  DocumentFieldTypeString,
					Value: "unique_id",
				},
				"name": {
					Type:  DocumentFieldTypeString,
					Value: "John Doe",
				},
			},
		})
		err := store.DumpToFile(tmpFile)
		assert.Nil(t, err)
		store2, err := NewStoreFromFile(tmpFile)
		assert.Nil(t, err)
		assert.Equal(t, store, store2, "Stores should be the same")
	})

	t.Run("Should return error on invalid file", func(t *testing.T) {
		_, err := NewStoreFromFile("nonexistent.json")
		assert.Error(t, err)
	})
}
