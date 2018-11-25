package api

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBookToJSON(t *testing.T) {
	book := Book{Title: "Cloud Native Go", Author: "Reimer", ISBN: "012345678"}
	json := book.ToJSON()

	assert.Equal(t, `{"title":"Cloud Native Go","author":"Reimer","isbn":"012345678"}`,
		string(json), "Book JSON marshalling wrong.")
}

func TestBookFromJSON(t *testing.T) {
	json := []byte(`{"title":"Cloud Native Go","author":"Reimer","isbn":"012345678"}`)
	book := FromJSON(json)

	assert.Equal(t, Book{Title: "Cloud Native Go", Author: "Reimer", ISBN: "012345678"},
		book, "Book JSON Marshalling wrong.")
}
