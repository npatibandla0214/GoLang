package main

import (
	"Assignment_4/app"
	"fmt"
)

func main() {
	//Add Accounts
	myAcc := app.Account{"Nandan", "Patibandla", 1, 10000}
	lstAccount := app.AddAccount(myAcc)
	fmt.Println(lstAccount)

	myAcc1 := app.Account{"Gagan", "Pabby", 2, 20000}
	lstAccount = app.AddAccount(myAcc1)
	fmt.Println(lstAccount)

	myAcc2 := app.Account{"Kushal", "Kush", 3, 30000}
	lstAccount = app.AddAccount(myAcc2)
	fmt.Println(lstAccount)

	//Delete Accounts
	app.RemoveAccount(1)
	fmt.Println(lstAccount)

	//Balance
	lstAccounts := app.ListAccounts()
	for k, _ := range lstAccounts {
		fmt.Printf("Balance for acct holder %d  : %d \n", lstAccounts[k].AccountNumber, app.AccountBalance(k))
	}

	//Deposit Money
	fmt.Printf("Balance before deposit: %d \n", lstAccounts[2].Balance)
	fmt.Printf("Balance after deposit: %d \n", app.DepositMoney(2, 500))

	//Withdraw Money
	fmt.Printf("Balance before deposit: %d \n", lstAccounts[3].Balance)
	fmt.Printf("Balance after deposit: %d \n", app.WithdrawMoney(3, 500))

	//Transfer Money
	fmt.Printf("Balance of the depositor after transfer: %d \n", lstAccounts[2].Balance)
	fmt.Printf("Balance of the depositee after transfer: %d \n", lstAccounts[3].Balance)
	depositorBalance, DepositeeBalance := app.TransferMoney(2, 3, 10000)
	fmt.Printf("Balance of the depositor after transfer: %d \n", depositorBalance)
	fmt.Printf("Balance of the depositee after transfer: %d \n", DepositeeBalance)
}
