package createtransaction

import (
	"testing"

	"github.com.br/PMafra/wallet-core/internal/entity"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type AccountGatewayMock struct {
	mock.Mock
}

func (m *AccountGatewayMock) FindByID(id string) (*entity.Account, error) {
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

func TestCreateTransactionUseCase_Execute(t *testing.T) {
	client1, _ := entity.NewClient("John Doe", "john@doe.com")
	accountFrom := entity.NewAccount(client1)
	accountFrom.Credit(1000.0)

	client2, _ := entity.NewClient("John Doe 2", "john2@doe.com")
	accountTo := entity.NewAccount(client2)
	accountTo.Credit(1000.0)

	accountGatewayMock := &AccountGatewayMock{}
	accountGatewayMock.On("FindByID", accountFrom.ID).Return(accountFrom, nil)
	accountGatewayMock.On("FindByID", accountTo.ID).Return(accountTo, nil)

	transactionGatewayMock := &TransactionGatewayMock{}
	transactionGatewayMock.On("Create", mock.Anything).Return(nil)

	uc := NewCreateTransactionUseCase(accountGatewayMock, transactionGatewayMock)
	inputDTO := &CreateTransactionInputDTO{
		AccountFromID: accountFrom.ID,
		AccountToID:   accountTo.ID,
		Amount:        100.0,
	}
	outputDTO, err := uc.Execute(inputDTO)
	assert.Nil(t, err)
	assert.NotNil(t, outputDTO.TransactionID)
	accountGatewayMock.AssertExpectations(t)
	accountGatewayMock.AssertNumberOfCalls(t, "FindByID", 2)
	transactionGatewayMock.AssertExpectations(t)
	transactionGatewayMock.AssertNumberOfCalls(t, "Create", 1)
}
