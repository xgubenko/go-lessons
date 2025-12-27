package pointers

import (
	"fmt"
)

type Bitcoin int

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}

type Wallet struct {
	balance Bitcoin
}

var ErrInsufficientFunds = fmt.Errorf("cannot withdraw, insufficient funds")

func (w *Wallet) Deposit(sum Bitcoin) {
	w.balance += sum
}

func (w Wallet) Balance() Bitcoin {
	return w.balance
}

func (w *Wallet) Withdraw(sum Bitcoin) error {
	if sum > w.balance {
		return ErrInsufficientFunds
	}
	w.balance -= sum
	return nil
}
