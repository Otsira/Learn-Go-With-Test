package pointers_errors

import (
	"errors"
	"fmt"
)

type Bitcoin int

var ErrInsuficcientFunds = errors.New("cannot withdraw, insufficient funds")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

func (wallet *Wallet) Deposit(amount Bitcoin) {
	wallet.balance += amount
}
func (wallet *Wallet) Withdraw(amount Bitcoin) error {
	if amount > wallet.balance {
		return ErrInsuficcientFunds
	}
	wallet.balance -= amount
	return nil
}
func (wallet *Wallet) Balance() Bitcoin {
	return wallet.balance
}
