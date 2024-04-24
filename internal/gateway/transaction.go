package gateway

import "github.com.br/PMafra/wallet-core/internal/entity"

type TransactionGateway interface {
	Create(transaction *entity.Transaction) error
}
