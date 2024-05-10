package pointers

import (
	"testing"
)

func TestWallet(t *testing.T) {
	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(Bitcoin(10))
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{10}
		err := wallet.Withdraw(Bitcoin(5))
		want := Bitcoin(5)

		assertError(t, err)
		assertBalance(t, wallet, want)
	})
}

func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}

func assertError(t testing.TB, got error) {
	t.Helper()
	if got != nil {
		t.Errorf("got error: %s", got.Error())
	}
}
