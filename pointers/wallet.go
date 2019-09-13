package pointers

import "fmt"

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
