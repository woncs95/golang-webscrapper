package accounts

import (
	"errors"
	"fmt"
)

//Account struct
type Account struct {
	owner   string
	balance int
}

var errNoMoney = errors.New("can't withdraw more than amount")

//NewAccount creates Account function creates struct
func NewAccount(owner string) *Account {
	account := Account{owner: owner, balance: 0}
	return &account
}

// Deposit some amount on your account
func (a *Account) Deposit(amount int) {
	a.balance += amount
}

// Balance of account
func (a Account) Balance() int {
	return a.balance
}

//Withdraw x from your account
func (a *Account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMoney
	}
	a.balance -= amount
	return nil
}

func (a *Account) ChangeOwner(newOwner string) {
	a.owner = newOwner

}

func (a Account) Owner() string {
	return a.owner
}

//This returns whatever I want for fmt.Println(account) as string
func (a Account) String() string {
	return fmt.Sprint(a.Owner(), "'s account \nhas: ", a.Balance())
}
