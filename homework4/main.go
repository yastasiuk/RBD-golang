package main

import (
	"fmt"
	. "homework4/documentstore"
)

func main() {
	store := NewStore()
	_, collection := store.CreateCollection("test", &CollectionConfig{PrimaryKey: "id"})
	collection.Put(Document{
		Fields: map[string]DocumentField{
			"id": {Type: DocumentFieldTypeString, Value: "unique_id"},
		},
	})

	document, _ := collection.Get("unique_id")

	fmt.Printf("Store: %+v", store)
	fmt.Println()
	fmt.Printf("Collection: %+v", collection)
	fmt.Println()
	fmt.Printf("Document: %+v", document)
	fmt.Println()
}
