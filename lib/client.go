package lib

import (
	"fmt"
	"time"

	rs "github.com/RediSearch/redisearch-go/redisearch"
)

const (
	HOST = "localhost:7000"
)

type Client = rs.Client
type Schema = rs.Schema
type Document = rs.Document

func Print(values ...interface{}) {
	fmt.Println(values...)
}

func Now() int64 {
	return time.Now().Unix()
}

// Create a client. By default a client is schemaless
// unless a schema is provided when creating the index
func GetClient(index string) *Client {
	return rs.NewClient(HOST, index)
}

func NewTextField(name string, weight float64, sortable bool) rs.Field {
	options := rs.TextFieldOptions{
		Weight:   float32(weight),
		Sortable: sortable,
	}
	return rs.NewTextFieldOptions(name, options)
}

func NewNumericField(name string) rs.Field {
	return rs.NewNumericField(name)
}

func NewSchema() *Schema {
	return rs.NewSchema(rs.DefaultOptions)
}

func NewDocument(id string) Document {
	weight := float32(1.0)
	return rs.NewDocument(id, weight)
}

func NewQuery(query string) *rs.Query {
	return rs.NewQuery(query)
}

func DropIndex(client *rs.Client) error {
	return client.Drop()
}

func CreateIndex(client *rs.Client, schema *rs.Schema) error {
	return client.CreateIndex(schema)
}

func IndexDoc(client *rs.Client, doc Document) error {
	return client.Index([]Document{doc}...)
}
