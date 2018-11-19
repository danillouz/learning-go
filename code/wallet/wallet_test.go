package wallet

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assert := func(t *testing.T, got, want Bitcoin) {
		t.Helper()

		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		value := Bitcoin(10.0)

		wallet.Deposit(value)

		got := wallet.Balance()
		want := value

		assert(t, got, want)
	})

	t.Run("Withdraw", func(t *testing.T) {
		start := Bitcoin(10.0)
		withdrawed := Bitcoin(1.0)
		wallet := Wallet{balance: start}

		wallet.Withdraw(withdrawed)

		got := wallet.Balance()
		want := start - withdrawed

		assert(t, got, want)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10.0)}

		err := wallet.Withdraw(Bitcoin(20.0))

		if err == nil {
			t.Fatalf("got %v, want error", err)
		}

		if err != InsufficientFundsError {
			t.Errorf("got error %v, want error %v", err, InsufficientFundsError)
		}
	})
}
