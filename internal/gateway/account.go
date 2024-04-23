package gateway

import "github.com.br/PMafra/wallet-core/internal/entity"

type AccountGateway interface {
	FindByID(id string) (*entity.Account, error)
	Save(account *entity.Account) error
}
