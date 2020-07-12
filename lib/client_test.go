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
	query := NewQuery("hello world")
	query.Limit(0, 2)
	query.SetReturnFields(FIELD_TITLE)
	return client.Search(query)
}

func TestSearchHelloWorld(t *testing.T) {
	schema := createExampleSchema()
	client := GetClient("testIndex")
	DropIndex(client)
	err := CreateIndex(client, schema)
	assert.Nil(t, err)

	doc := createExampleDoc("doc1", "hello world", "this is the final countdown")
	err = IndexDoc(client, doc)
	assert.Nil(t, err)

	docs, total, err := searchHelloWorld(client)
	Print(err)
	Print(docs)
	Print(total)

	//fmt.Println(docs[0].Id, docs[0].Properties["title"], total, err)
	// Output: doc1 Hello world 1 <nil>
}
