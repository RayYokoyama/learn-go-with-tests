package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertError := func(t *testing.T, err error, want string) {
		t.Helper()
		if err == nil {
			t.Error("wanted an error but didn't get one")
		}

		if err.Error() != want {
			t.Errorf("got %q, want %q", err, want)
		}
	}

	assertNoError := func(t *testing.T, err error) {
		t.Helper()
		if err != nil {
			t.Error("got an error but didn't want one")
		}
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}
		err := wallet.Withdraw(Bitcoin(12))
		want := Bitcoin(8)
		assertBalance(t, wallet, want)
		assertNoError(t, err)
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)
		wallet := Wallet{startingBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, startingBalance)
		assertError(t, err, ErrInsufficientFunds.Error())
	})
}
