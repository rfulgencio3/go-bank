package main

import (
	a "_projects/go-bank/accounts"
	c "_projects/go-bank/clients"
	"fmt"
)

func main() {

	fmt.Println("CheckingAccount")
	personData := c.Client{"Ricardo", "system analyst"}
	checkingAccount := a.CheckingAccount{
		Client:        personData,
		Agency:        "001",
		AccountNumber: 123456,
	}

	fmt.Println(checkingAccount)

	status, value := checkingAccount.Deposit(200)
	fmt.Println(status, "Deposit value:", value)

	withdral, balance := checkingAccount.Withdraw(25.5), checkingAccount.GetBalance()
	fmt.Println(withdral, "Actual balance:", balance)

	PayBillet(&checkingAccount, 50)
	fmt.Println("Final balance:", checkingAccount.GetBalance())

	fmt.Println("")

	fmt.Println("SavingsAccount")
	savingsAccount := a.SavingsAccount{
		Client:        personData,
		Agency:        "001",
		AccountNumber: 123456,
		Operation:     1,
	}

	fmt.Println(savingsAccount)

	savingAccountStatus, savingAccountValue := savingsAccount.Deposit(200)
	fmt.Println(savingAccountStatus, "Deposit value:", savingAccountValue)

	savingsAccountWithdraw, savingsAccountBalance := savingsAccount.Withdraw(25.5), savingsAccount.GetBalance()
	fmt.Println(savingsAccountWithdraw, "Actual balance:", savingsAccountBalance)
}

func PayBillet(account verifyAccount, billetAmount float64) {
	account.Debit(billetAmount)
}

type verifyAccount interface {
	Debit(amount float64) string
}
