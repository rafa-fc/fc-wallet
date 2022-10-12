package gateway

import "github.com/rafassts/fullcycle/9-projeto-wallet/walletcore/internal/entity"

type ClientGateway interface {
	Get(id string) (*entity.Client, error)
	Save(client *entity.Client) error
}
