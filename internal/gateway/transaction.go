package gateway

import "github.com/aluiziodeveloper/ms-wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
