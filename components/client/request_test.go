package client

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTransObject2Queries(t *testing.T) {
	age := 10

	q, err := transObject2Queries(struct {
		Name string `json:"name"`
		Age  *int   `json:"int"`
	}{
		Name: "alex",
		Age:  &age,
	})
	assert.Nil(t, err)
	assert.Equal(t, "int=10&name=alex", q.Encode())

	q, err = transObject2Queries(struct {
		Name string `json:"name"`
		Age  *int   `json:"int"`
	}{
		Name: "alex",
	})
	assert.Nil(t, err)
	assert.Equal(t, "name=alex", q.Encode())

	q, err = transObject2Queries(struct {
		Name string `json:"name"`
		Age  *int   `json:"int"`
	}{})
	assert.Nil(t, err)
	assert.Equal(t, "name=", q.Encode())
}
