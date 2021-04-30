package main

import (
	"fmt"

	"github.com/DongYeon/learnGo/account"
)

func main() {
	MyAccount := account.NewAccount("dongsik")
	fmt.Println(*MyAccount)
	MyAccount.Deposit(10000)
	MyAccount.Withdraw(6500)
	MyAccount.Withdraw(5500)
	fmt.Println(*MyAccount)
}
