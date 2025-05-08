package main

import (
	"fmt"
	. "lesson_07/document_store"
)

func init() {
	setupLogger()
}

func main() {
	store := NewStore()
	_, collection := store.CreateCollection("test", &CollectionConfig{PrimaryKey: "id"})
	collection.Put(Document{
		Fields: map[string]DocumentField{
			"id": {Type: DocumentFieldTypeString, Value: "unique_id"},
		},
	})

	if err := store.DumpToFile("test.json"); err != nil {
		fmt.Println("Error:", err)
	}

	store2, _ := NewStoreFromFile("test.json")

	collection2, _ := store2.GetCollection("test")
	document2, _ := collection2.Get("unique_id")
	fmt.Println("document2:", document2)
}
