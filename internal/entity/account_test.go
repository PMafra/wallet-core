package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewAccount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account := NewAccount(client)
	assert.Equal(t, client.ID, account.Client.ID)
	assert.Equal(t, 0.0, account.Balance)
}

func TestCreateNewAccountWithBlankClient(t *testing.T) {
	account := NewAccount(nil)
	assert.Nil(t, account)
}

func TestCreditAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account := NewAccount(client)
	account.Credit(100.0)
	assert.Equal(t, float64(100), account.Balance)
}

func TestDebitAmount(t *testing.T) {
	client, _ := NewClient("John Doe", "john@doe.com")
	account := NewAccount(client)
	account.Debit(100.0)
	assert.Equal(t, float64(-100), account.Balance)
}
