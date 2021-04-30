package account

import (
	"errors"
	"fmt"
)

type account struct {
	owner   string
	balance int
}

var errNoMey = errors.New("You dont have enough money")

func NewAccount(name string) *account {
	NewAccount := account{owner: name, balance: 0}
	return &NewAccount
}

func (a *account) Deposit(amount int) {
	a.balance += amount
	fmt.Println(a.balance, "deposited")
}

func (a *account) Withdraw(amount int) error {
	if a.balance < amount {
		return errNoMey
	}
	a.balance -= amount
	fmt.Println(a.balance, "withdrawn")
	return nil
}

func (a *account) ChangeOwner(newOwner string) {
	a.owner = newOwner
}

func (a account) String() string {
	return fmt.Sprint("---------------------\n",
		"Owner : ", a.owner, "\nbalance : ", a.balance,
		"\n---------------------")
}

func (a account) CheckMoeny() {
	fmt.Println(a.balance)
}
