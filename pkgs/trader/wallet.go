package trader

import (
	"errors"
	"github.com/shopspring/decimal"
)

var (
	ErrInsufficientFunds = errors.New("cannot draw, insufficient funds")
)

type Wallet struct {
	balance decimal.Decimal
}

func (w *Wallet) Deposit(money decimal.Decimal) {
	w.balance = w.balance.Add(money)
}

func (w *Wallet) Withdraw(money decimal.Decimal) error {
	result := w.balance.Sub(money)
	if result.IsNegative() {
		return ErrInsufficientFunds
	}

	w.balance = result
	return nil
}

func (w *Wallet) Balance() decimal.Decimal {
	return w.balance
}
