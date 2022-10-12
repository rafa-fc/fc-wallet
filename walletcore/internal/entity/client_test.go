package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewClient(t *testing.T) {
	client, err := NewClient("John Doe", "joe@example.com")
	assert.Nil(t, err)
	assert.NotNil(t, client)
	assert.Equal(t, "John Doe", client.Name)
	assert.Equal(t, "joe@example.com", client.Email)
}

func TestCreateNewClientWhenInvalid(t *testing.T) {
	client, err := NewClient("", "")
	assert.NotNil(t, err)
	assert.Nil(t, client)
}

func TestUpdateClient(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@example.com")
	err := client.Update("John Doe 2", "joe2@example.com")
	assert.Nil(t, err)
	assert.Equal(t, "John Doe 2", client.Name)
	assert.Equal(t, "joe2@example.com", client.Email)

}

func TestUpdateClientWhenInvalid(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@example.com")

	err := client.Update("", "")
	assert.Error(t, err, "name is required")

	err = client.Update("John", "")
	assert.Error(t, err, "email is required")

}

func TestAddAccountToClient(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@example.com")
	account := NewAccount(client)
	err := client.AddAccount(account)

	assert.Nil(t, err)
	assert.Equal(t, 1, len(client.Accounts))

}
