package pointers

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

//
func TestWallet(t *testing.T) {
	wallet := Wallet{}

	wallet.Deposit(Bitcoin(10))
	fmt.Printf("address of balance in test is %v \n", &wallet.balance)

	got := wallet.Balance()
	want := Bitcoin(10)

	assert.Equal(t, want, got)
}
