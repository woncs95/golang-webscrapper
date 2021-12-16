package banking

import (
	"accounts"
	"fmt"
)

func Banking() {
	account := accounts.NewAccount("chang")
	account.Deposit(10)
	fmt.Println(account)
	// err := account.Withdraw(30)
	// if err != nil {
	// 	//or this : fmt.Println(err)
	// 	fmt.Println(err)
	// }
	// fmt.Println(account.Balance(), account.Owner())
	// account.ChangeOwner("amy")
	// fmt.Println(account.Owner())
}
