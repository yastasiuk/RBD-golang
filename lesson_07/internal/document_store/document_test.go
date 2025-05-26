package documentstore

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestGetField_StringValue(t *testing.T) {
	doc := Document{
		Fields: map[string]DocumentField{
			"name": {
				Type:  DocumentFieldTypeString,
				Value: "123456",
			},
		},
	}

	value := doc.GetField("name")
	assert.Equal(t, "123456", value)
}
