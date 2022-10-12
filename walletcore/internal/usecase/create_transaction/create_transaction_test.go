package createtransaction

import (
	"testing"

	"github.com/rafassts/fullcycle/9-projeto-wallet/walletcore/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindById(id string) (*entity.Account, error) {
	args := m.Called(id)
	return args.Get(0).(*entity.Account), args.Error(1)
}

func (m *AccountGatewayMock) Save(account *entity.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

type TransactionGatewayMock struct {
	mock.Mock
}

func (m *TransactionGatewayMock) Create(transaction *entity.Transaction) error {
	args := m.Called(transaction)
	return args.Error(0)
}

func TestCreateTransactionUseCase(t *testing.T) {
	client1, _ := entity.NewClient("Rafa", "rafa@example.com")
	client2, _ := entity.NewClient("Suzi", "suzie@example.com")
	account1 := entity.NewAccount(client1)
	account1.Credit(1000)
	account2 := entity.NewAccount(client2)
	account2.Credit(500)

	mockAccount := &AccountGatewayMock{}
	mockAccount.On("FindById", account1.Id).Return(account1, nil)
	mockAccount.On("FindById", account2.Id).Return(account2, nil)

	mockTransaction := &TransactionGatewayMock{}
	mockTransaction.On("Create", mock.Anything).Return(nil)

	inputDto := CreateTransactionInputDto{
		AccountIdFrom: account1.Id,
		AccountIdTo:   account2.Id,
		Amount:        100,
	}

	uc := NewCreateTransactionUseCase(mockTransaction, mockAccount)

	output, err := uc.Execute(inputDto)

	assert.Nil(t, err)
	assert.NotNil(t, output)
	mockAccount.AssertExpectations(t)
	mockTransaction.AssertExpectations(t)
	mockAccount.AssertNumberOfCalls(t, "FindById", 2)
	mockTransaction.AssertNumberOfCalls(t, "Create", 1)
}
