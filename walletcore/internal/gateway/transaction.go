package gateway

import "github.com/rafassts/fullcycle/9-projeto-wallet/walletcore/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
