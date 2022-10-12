package createaccount

import (
	"github.com/rafassts/fullcycle/9-projeto-wallet/walletcore/internal/entity"
	"github.com/rafassts/fullcycle/9-projeto-wallet/walletcore/internal/gateway"
)

type CreateAccountInputDto struct {
	ClientId string
}

type CreateAccountOutputDto struct {
	Id string
}

// variaveis da injeção de dependencia
type CreateAccountUseCase struct {
	AccountGateway gateway.AccountGateway
	ClientGateway  gateway.ClientGateway
}

// construtor
func NewCreateAccountUseCase(ag gateway.AccountGateway, cg gateway.ClientGateway) *CreateAccountUseCase {
	return &CreateAccountUseCase{
		AccountGateway: ag,
		ClientGateway:  cg,
	}
}

func (uc *CreateAccountUseCase) Execute(input CreateAccountInputDto) (*CreateAccountOutputDto, error) {
	client, err := uc.ClientGateway.Get(input.ClientId)

	if err != nil {
		return nil, err
	}

	account := entity.NewAccount(client)

	err = uc.AccountGateway.Save(account)

	if err != nil {
		return nil, err
	}

	return &CreateAccountOutputDto{
		Id: account.Id,
	}, nil

}
