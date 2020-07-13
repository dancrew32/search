package lib

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

const (
	FIELD_BODY  = "body"
	FIELD_TITLE = "title"
	FIELD_DATE  = "date"
)

func createExampleSchema() *Schema {
	title := NewTextField(FIELD_TITLE, 5.0, true)
	body := NewTextField(FIELD_BODY, 1.0, false)
	date := NewNumericField(FIELD_DATE)
	schema := NewSchema()
	schema.AddField(title)
	schema.AddField(body)
	schema.AddField(date)
	return schema
}

func createExampleDoc(id, title, body string) Document {
	now := Now()
	doc := NewDocument(id)
	doc.Set(FIELD_TITLE, title)
	doc.Set(FIELD_BODY, body)
	doc.Set(FIELD_DATE, now)
	return doc
}

func searchHelloWorld(client *Client) ([]Document, int, error) {
	query := NewQuery("hello")
	query.Limit(0, 4)
	query.SetReturnFields(FIELD_TITLE, FIELD_BODY)
	return client.Search(query)
}

func TestSearchHelloWorld(t *testing.T) {
	schema := createExampleSchema()
	client := GetClient("testIndex")
	DropIndex(client)
	err := CreateIndex(client, schema)
	assert.Nil(t, err)

	doc1 := createExampleDoc("doc1", "hello world", "this is the final countdown")
	err = IndexDoc(client, doc1)
	assert.Nil(t, err)

	doc2 := createExampleDoc("doc2", "world hello", "beans")
	err = IndexDoc(client, doc2)
	assert.Nil(t, err)

	// Unmatched
	doc3 := createExampleDoc("doc3", "unrelated", "document")
	err = IndexDoc(client, doc3)
	assert.Nil(t, err)

	doc4 := createExampleDoc("doc4", "unrelated title", "hello")
	err = IndexDoc(client, doc4)
	assert.Nil(t, err)

	docs, total, err := searchHelloWorld(client)
	assert.Nil(t, err)
	assert.Equal(t, 3, total)
	assert.Equal(t, "beans", docs[0].Properties[FIELD_BODY])
	assert.Equal(t, "this is the final countdown", docs[1].Properties[FIELD_BODY])
	assert.Equal(t, "hello", docs[2].Properties[FIELD_BODY])
}
