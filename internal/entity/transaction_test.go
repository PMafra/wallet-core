package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateNewTransaction(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@doe.com")
	client2, _ := NewClient("John Doe 2", "john@doe2.com")

	account1 := NewAccount(client1)
	account2 := NewAccount(client2)

	account1.Credit(1000.0)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, 100.0)
	assert.NotNil(t, transaction)
	assert.Nil(t, err)
	assert.Equal(t, 900.0, account1.Balance)
	assert.Equal(t, 1100.0, account2.Balance)
}

func TestCreateTransactionWithInsuficientBalance(t *testing.T) {
	client1, _ := NewClient("John Doe", "john@doe.com")
	client2, _ := NewClient("John Doe 2", "john@doe2.com")

	account1 := NewAccount(client1)
	account2 := NewAccount(client2)

	account1.Credit(1000.0)
	account2.Credit(1000.0)

	transaction, err := NewTransaction(account1, account2, 2000.0)
	assert.Nil(t, transaction)
	assert.NotNil(t, err)
	assert.Error(t, err, "insufficient balance")

	assert.Equal(t, 1000.0, account1.Balance)
	assert.Equal(t, 1000.0, account2.Balance)
}
