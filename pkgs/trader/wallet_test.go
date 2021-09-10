package trader

import (
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("can add and remove money from the bank", func(t *testing.T) {
		wallet := Wallet{}
		valueToDeposit, err := decimal.NewFromString("1000")
		assert.NoError(t, err)
		wallet.Deposit(valueToDeposit)
		valueToWithdraw, err := decimal.NewFromString("350")
		assert.NoError(t, err)
		err = wallet.Withdraw(valueToWithdraw)
		assert.NoError(t, err)
		expected, err := decimal.NewFromString("650")
		assert.NoError(t, err)
		assert.Equal(t, expected, wallet.Balance())
	})

	t.Run("cannot remove more money than we have", func(t *testing.T) {
		wallet := Wallet{}
		valueToDeposit, err := decimal.NewFromString("350")
		assert.NoError(t, err)
		wallet.Deposit(valueToDeposit)
		valueToWithdraw, err := decimal.NewFromString("2000")
		assert.NoError(t, err)
		err = wallet.Withdraw(valueToWithdraw)
		assert.EqualError(t, err, ErrInsufficientFunds.Error())
		expected, err := decimal.NewFromString("350")
		assert.NoError(t, err)
		assert.Equal(t, expected, wallet.Balance())
	})

}
