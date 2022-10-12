package gateway

import "github.com/rafassts/fullcycle/9-projeto-wallet/walletcore/internal/entity"

type AccountGateway interface {
	Save(account *entity.Account) error
	FindById(id string) (*entity.Account, error)
}
