package main

import (
	"fmt"
	. "homework3/documentstore"
)

func main() {
	doc1 := Document{}
	doc1.Fields = map[string]DocumentField{}
	doc1.Fields["key"] = DocumentField{
		Type:  DocumentFieldTypeString,
		Value: "value",
	}
	Put(doc1)
	list := List()
	fmt.Println(list[0])

}

/**
Add benchmark with and w/o reflection(for type validation)
*/
