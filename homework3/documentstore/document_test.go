package documentstore

import (
	"fmt"
	"reflect"
	"testing"
)

func TestListEmpty(t *testing.T) {
	defer ResetStore()

	t.Run("Should insert new document", func(t *testing.T) {
		ans := List()
		if len(ans) != 0 {
			t.Error(fmt.Sprintf("Empty store should return 0 elements. Returned: %d", len(ans)))
		}
	})
}

func TestPutNew(t *testing.T) {
	defer ResetStore()

	doc := Document{}
	doc.Fields = map[string]DocumentField{
		"key":  {Value: "123", Type: DocumentFieldTypeString},
		"name": {Value: "Joe", Type: DocumentFieldTypeString},
	}

	t.Run("Should insert new document", func(t *testing.T) {
		Put(doc)
		if ans := List(); len(ans) != 1 {
			t.Error(fmt.Sprintf("Should return 1 document, returned %d:", len(ans)))
		} else if !reflect.DeepEqual(ans[0], doc) {
			t.Error("New document is not equal to an old one")
		}
	})
}

func TestPutInvalid(t *testing.T) {
	defer ResetStore()

	testCases := []Document{
		{Fields: map[string]DocumentField{
			"key": {Value: "Not valid type", Type: DocumentFieldTypeBool},
		}},
		{Fields: map[string]DocumentField{
			"notKey": {Value: "Key field is missing", Type: DocumentFieldTypeString},
		}},
		{Fields: map[string]DocumentField{
			"key": {Value: 123, Type: DocumentFieldTypeString}, // Invalid value type
		}},
	}

	for i, testCase := range testCases {
		ResetStore()
		t.Run(fmt.Sprintf("Test case #%d", i), func(t *testing.T) {
			Put(testCase)
			ans := List()
			if len(ans) != 0 {
				t.Error(fmt.Sprintf("Test case #%d added new item to documents store. Size: %d", i, len(ans)))
			}
		})
	}
}

func TestPutOverride(t *testing.T) {
	defer ResetStore()

	key := "not_unique_key"

	doc1 := Document{Fields: map[string]DocumentField{
		"key":        {Value: key, Type: DocumentFieldTypeString},
		"extraField": {Value: "should be removed", Type: DocumentFieldTypeString},
	}}

	doc2 := Document{Fields: map[string]DocumentField{
		"key":      {Value: key, Type: DocumentFieldTypeString},
		"newField": {Value: "new value", Type: DocumentFieldTypeString},
	}}

	t.Run(fmt.Sprintf("Should override already existing document in store if key overlap"), func(t *testing.T) {
		Put(doc1)
		Put(doc2)

		list := List()
		if len(list) != 1 {
			t.Error(fmt.Sprintf("Store should have only 1 document, got: %d", len(list)))
		} else if !reflect.DeepEqual(list[0], doc2) {
			t.Error(fmt.Sprintf("Store document is not equal to 'doc2', got: %#v", list[0]))
		}
	})
}

func TestGetNonExistingDoc(t *testing.T) {
	defer ResetStore()

	t.Run(fmt.Sprintf("Should return nil on GET non existing key"), func(t *testing.T) {
		if ans, ok := Get("not_existing_key"); ok {
			t.Error("Store returned 'true' on GET(..)")
		} else if ans != nil {
			t.Error(fmt.Sprintf("Store returned non nil value. Got: %#v", ans))
		}
	})
}

func TestGetExistingDoc(t *testing.T) {
	defer ResetStore()

	key := "existing_key"
	doc := Document{Fields: map[string]DocumentField{
		"key": {Value: key, Type: DocumentFieldTypeString},
	}}

	t.Run(fmt.Sprintf("Should return document from store"), func(t *testing.T) {
		Put(doc)
		ans, ok := Get(key)
		if !ok {
			t.Error("Store returned 'false' on GET(..)")
		}
		if !reflect.DeepEqual(*ans, doc) {
			t.Error(fmt.Sprintf("Store returned wrong document. Got: %#v", ans))
		}
	})
}

func TestDeleteExistingKey(t *testing.T) {
	defer ResetStore()
	key := "existing_key"
	doc := Document{Fields: map[string]DocumentField{
		"key": {Value: key, Type: DocumentFieldTypeString},
	}}

	t.Run(fmt.Sprintf("Should correctly delete document"), func(t *testing.T) {
		Put(doc)
		deleted := Delete(key)
		if !deleted {
			t.Error("Store returned 'false' on deleting valid key")
		}
		if docFromStore, _ := Get(key); docFromStore != nil {
			t.Error(fmt.Sprintf("Documet was not removed from store %#v", docFromStore))
		}
	})
}
