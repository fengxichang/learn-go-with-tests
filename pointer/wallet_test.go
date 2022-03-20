package pointer

import (
	"testing"
)

func TestWallet(t *testing.T) {
	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		got := wallet.Balance()

		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

	assertNotError := func(t *testing.T, got error) {
		t.Helper()
		if got != nil {
			t.Fatal("got an error but didn't want one")
		}
	}

	assertError := func(t *testing.T, got error, want error) {
		t.Helper()

		if got == nil {
			t.Fatal("wanted an error but didn't get one")
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
	})

	t.Run("withdraw", func(t *testing.T) {
		wallet := Wallet{balance: Bitcoin(20)}

		err := wallet.Withdraw(Bitcoin(10))

		want := Bitcoin(10)

		assertBalance(t, wallet, want)
		assertNotError(t, err)
	})

	t.Run("withdraw insufficient funds", func(t *testing.T) {
		startingBalance := Bitcoin(20)

		wallet := Wallet{startingBalance}

		err := wallet.Withdraw(Bitcoin(100))

		assertBalance(t, wallet, startingBalance)

		assertError(t, err, ErrInsufficientFunds)
	})
}
