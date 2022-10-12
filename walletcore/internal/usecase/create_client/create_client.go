package createclient

import (
	"time"

	"github.com/rafassts/fullcycle/9-projeto-wallet/walletcore/internal/entity"
	"github.com/rafassts/fullcycle/9-projeto-wallet/walletcore/internal/gateway"
)

type CreateClientInputDto struct {
	Name  string
	Email string
}

type CreateClientOutputDto struct {
	Id        string
	Name      string
	Email     string
	CreatedAt time.Time
	UpdatedAt time.Time
}

// variaveis da injeção de dependencia
type CreateClientUseCase struct {
	ClientGateway gateway.ClientGateway
}

// construtor para createclientusecase
func NewCreateClientUseCase(clientGateway gateway.ClientGateway) *CreateClientUseCase {
	return &CreateClientUseCase{
		ClientGateway: clientGateway,
	}
}

func (uc *CreateClientUseCase) Execute(input CreateClientInputDto) (*CreateClientOutputDto, error) {
	client, err := entity.NewClient(input.Name, input.Email)

	if err != nil {
		return nil, err
	}

	err = uc.ClientGateway.Save(client)

	if err != nil {
		return nil, err
	}

	return &CreateClientOutputDto{
		Id:        client.Id,
		Name:      client.Name,
		Email:     client.Email,
		CreatedAt: client.CreatedAt,
		UpdatedAt: client.UpdatedAt,
	}, nil

}
