package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateTransaction(t *testing.T) {
	client1, _ := NewClient("Rafa", "rafa@example.com")
	client2, _ := NewClient("Suzi", "suzie@example.com")
	account1 := NewAccount(client1)
	account1.Credit(1000)
	account2 := NewAccount(client2)
	account2.Credit(500)

	transaction, err := NewTransaction(account1, account2, 300)

	assert.Nil(t, err)
	assert.NotNil(t, transaction)
	assert.Equal(t, account1.Balance, float64(700))
	assert.Equal(t, account2.Balance, float64(800))

}

func TestCreateTransactionNotEnoughBalance(t *testing.T) {
	client1, _ := NewClient("Rafa", "rafa@example.com")
	client2, _ := NewClient("Suzi", "suzie@example.com")
	account1 := NewAccount(client1)
	account1.Credit(100)
	account2 := NewAccount(client2)
	account2.Credit(500)

	transaction, err := NewTransaction(account1, account2, 300)

	assert.NotNil(t, err)
	assert.Nil(t, transaction)
	assert.Error(t, err, "insufficient balance")
	assert.Equal(t, account1.Balance, float64(100))
	assert.Equal(t, account2.Balance, float64(500))

}
