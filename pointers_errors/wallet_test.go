package pointers_errors

import (
	"testing"
)

func TestWallet(t *testing.T) {

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{balance: 20}
		wallet.Withdraw(Bitcoin(10))
		want := Bitcoin(10)
		assertBalance(t, wallet, want)
	})
	t.Run("Withdraw more money that what's on the wallet", func(t *testing.T) {
		initBalance := Bitcoin(20)
		wallet := Wallet{initBalance}
		err := wallet.Withdraw(Bitcoin(100))
		assertBalance(t, wallet, initBalance)
		assertError(t, err, ErrInsuficcientFunds)
	})

}
func assertBalance(t testing.TB, wallet Wallet, want Bitcoin) {
	t.Helper()
	got := wallet.Balance()
	if got != want {
		t.Errorf("got %s want %s", got, want)
	}
}
func assertError(t testing.TB, got error, want error) {
	t.Helper()
	if got == nil {
		t.Fatal("i should have an error here")
	}
	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
