package entity

import (
	"errors"
	"time"

	"github.com/google/uuid"
)

type Transaction struct {
	Id          string
	AccountFrom *Account
	AccountTo   *Account
	Amount      float64
	CreatedAt   time.Time
}

func NewTransaction(accountFrom *Account, accountTo *Account, amount float64) (*Transaction, error) {
	transaction := &Transaction{
		Id:          uuid.New().String(),
		AccountFrom: accountFrom,
		AccountTo:   accountTo,
		Amount:      amount,
		CreatedAt:   time.Now(),
	}

	err := transaction.Validate()

	if err != nil {
		return nil, err
	}

	transaction.Transfer()

	return transaction, nil
}

func (t *Transaction) Transfer() {
	t.AccountFrom.Debit(t.Amount)
	t.AccountTo.Credit(t.Amount)
}

func (t *Transaction) Validate() error {
	if t.AccountFrom == nil {
		return errors.New("source account invalid")
	}

	if t.AccountTo == nil {
		return errors.New("target account invalid")
	}
	if t.Amount <= 0 {
		return errors.New("amount must be greater than zero")
	}
	if t.AccountFrom.Balance < t.Amount {
		return errors.New("insufficient balance")
	}

	return nil
}
