package createtransaction

import (
	"github.com.br/PMafra/wallet-core/internal/entity"
	"github.com.br/PMafra/wallet-core/internal/gateway"
)

type CreateTransactionInputDTO struct {
	AccountFromID string
	AccountToID   string
	Amount        float64
}

type CreateTransactionOutputDTO struct {
	TransactionID string
}

type CreateTransactionUseCase struct {
	AccountGateway     gateway.AccountGateway
	TransactionGateway gateway.TransactionGateway
}

func NewCreateTransactionUseCase(accountGateway gateway.AccountGateway, transactionGateway gateway.TransactionGateway) *CreateTransactionUseCase {
	return &CreateTransactionUseCase{
		AccountGateway:     accountGateway,
		TransactionGateway: transactionGateway,
	}
}

func (uc *CreateTransactionUseCase) Execute(input *CreateTransactionInputDTO) (*CreateTransactionOutputDTO, error) {
	accountFrom, err := uc.AccountGateway.FindByID(input.AccountFromID)
	if err != nil {
		return nil, err
	}
	accountTo, err := uc.AccountGateway.FindByID(input.AccountToID)
	if err != nil {
		return nil, err
	}
	transaction, err := entity.NewTransaction(accountFrom, accountTo, input.Amount)
	if err != nil {
		return nil, err
	}
	err = uc.TransactionGateway.Create(transaction)
	if err != nil {
		return nil, err
	}
	return &CreateTransactionOutputDTO{
		TransactionID: transaction.ID,
	}, nil
}
