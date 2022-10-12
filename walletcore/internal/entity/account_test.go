package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateAccounnt(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@example.com")
	account := NewAccount(client)
	assert.NotNil(t, account)
	assert.Equal(t, client.Id, account.Client.Id)
}

func TestCreateAccountWithoutClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@example.com")
	account := NewAccount(client)
	account.Credit(1000)
	assert.Equal(t, float64(1000), account.Balance)
}

func TestDebitAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "joe@example.com")
	account := NewAccount(client)
	account.Credit(1000)
	account.Debit(600)
	assert.Equal(t, float64(400), account.Balance)
}
