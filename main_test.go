package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestCreateContact(t *testing.T) {
	c := require.New(t)

	contact := NewContact("jhon", "555-55-55", "st. 123")

	c.Equal("jhon", contact.Name)
}
