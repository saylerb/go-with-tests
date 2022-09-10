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

		assertBalance(t, got, want)
	})
	t.Run("withdraw money from bitcoin wallet", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		if err != nil {
			t.Fatalf("did not expect an error to be returned, but one was: %v", err)

		}

		got := wallet.Balance()
		want := Bitcoin(10)

		assertBalance(t, got, want)
	})
	t.Run("overdrafting balance should result in an error", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(100))

		got := wallet.Balance()
		want := Bitcoin(20)

		assertBalance(t, got, want)

		if err == nil {
			// t.Fatal will stop the test, so that we don't assert on
			// error that does not exist
			t.Fatal("wanted an error to be retured, but no error was returned")
		}
		if err != ErrInsufficientBalance {
			t.Errorf("got %q, wanted %q", err, ErrInsufficientBalance)
		}
	})
}

func assertBalance(t *testing.T, got Bitcoin, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
