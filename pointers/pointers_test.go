package main

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("deposit money into bitcoin wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(10)}

		wallet.Deposit(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(20)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("withdraw money from bitcoin wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		wallet.Withdraw(Bitcoin(10))

		got := wallet.Balance()
		want := Bitcoin(10)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	})
	t.Run("overdrafting balance should result in an error", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		got := wallet.Balance()
		want := Bitcoin(20)

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}

		if err == nil {
			t.Error("wanted an error to be retured")
		}
	})
}