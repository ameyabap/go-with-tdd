package pointers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

//
func TestWallet(t *testing.T) {

	assertBalance := func(t *testing.T, wallet Wallet, want Bitcoin) {
		assert.Equal(t, want, wallet.Balance())
	}
	assertError := func(t *testing.T, want error, got error) {
		//assert.Equal(t, want, got)
		assert.Equal(t, want, got)
	}

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		err := wallet.Withdraw(Bitcoin(10))
		assertBalance(t, wallet, Bitcoin(10))
		assertError(t, err, nil)
	})

	t.Run("insufficient balance", func(t *testing.T) {
		wallet := Wallet{
			balance: Bitcoin(20),
		}

		err := wallet.Withdraw(Bitcoin(30))
		assertBalance(t, wallet, Bitcoin(20))
		assertError(t, err, ErrInsufficientBalance)
	})

}
