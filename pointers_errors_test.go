package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertIntMessage := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %d want %d", got, want)
		}
	}
	assertError := func(t *testing.T, got error, want string) {
		t.Helper()
		if got == nil {
			t.Errorf("wanted an error but didn't get one")
		}
		if got.Error() != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10)

		want := Bitcoin(10)
		assertIntMessage(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(15)
		_ = wallet.Withdraw(6)
		want := Bitcoin(9)
		assertIntMessage(t, wallet, want)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))

		assertError(t, err, "cannot withdraw, unsufficient funds")
		assertIntMessage(t, wallet, startingBalance)
	})
}
