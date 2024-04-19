package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "john@doe.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "john@doe.com", client.Email)
}

func TestCreateNewClientWhenArgsAreInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	err := client.Update("New John Doe", "new@doe.com")
	assert.Nil(t, err)
	assert.Equal(t, "New John Doe", client.Name)
	assert.Equal(t, "new@doe.com", client.Email)
}

func TestUpdateClientWithArgsAreInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	err := client.Update("", "")
	assert.Error(t, err, "name is required")
}
