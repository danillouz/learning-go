package wallet

import (
	"errors"
	"fmt"
)

var InsufficientFundsError = errors.New("Insufficient Funds")

// Bitcoin represents a cryptocurrency value.
type Bitcoin float64

func (b Bitcoin) String() string {
	return fmt.Sprintf("%.2f BTC", b)
}

// Wallet contains a balance and methods to interact with it.
type Wallet struct {
	balance Bitcoin
}

// Deposit will update the balance of the wallet by adding the provided value to it.
func (w *Wallet) Deposit(value Bitcoin) {
	w.balance += value
}

// Withdraw will update the balance of the wallet by removing the provided value from it.
func (w *Wallet) Withdraw(value Bitcoin) error {
	if value > w.balance {
		return InsufficientFundsError
	}

	w.balance -= value

	return nil
}

// Balance returns the current balance of the wallet.
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}
