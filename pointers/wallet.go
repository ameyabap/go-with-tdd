package pointers

import (
	"errors"
	"fmt"
)

// Bitcoin ...
type Bitcoin int

// Wallet ...
type Wallet struct {
	balance Bitcoin
}

// Deposit ...
func (w *Wallet) Deposit(amt Bitcoin) {
	fmt.Printf("address of balance in test is %v \n", &w.balance)
	w.balance += amt
}

// Balance ...
func (w *Wallet) Balance() Bitcoin {
	return w.balance
}

// Withdraw ...
func (w *Wallet) Withdraw(amt Bitcoin) error {
	if w.balance < amt {
		return ErrInsufficientBalance
	}
	w.balance -= amt
	return nil
}

// ErrInsufficientBalance ...
var ErrInsufficientBalance = errors.New("Can't Withdraw, Insufficient Balance")

func (b Bitcoin) String() string {
	return fmt.Sprintf("%d BTC", b)
}
