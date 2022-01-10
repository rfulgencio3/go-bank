package accounts

import c "_projects/go-bank/clients"

type CheckingAccount struct {
	Client        c.Client
	Agency        string
	AccountNumber int
	balance       float64
}

func (ca *CheckingAccount) Withdraw(amount float64) string {
	allowWithdrawal := amount > 0 && amount <= ca.balance
	if allowWithdrawal {
		ca.balance -= amount
		return "Withdrawal operation was successful."
	} else {
		return "Insuficient founds."
	}
}

func (ca *CheckingAccount) Debit(amount float64) string {
	allowWithdrawal := amount > 0 && amount <= ca.balance
	if allowWithdrawal {
		ca.balance -= amount
		return "Debit operation was successful."
	} else {
		return "Insuficient founds."
	}
}

func (ca *CheckingAccount) Deposit(amount float64) (string, float64) {
	if amount >= 0 {
		ca.balance += amount
		return "Deposit operation was successful.", ca.balance
	} else {
		return "Invalid value", ca.balance
	}
}

func (ca *CheckingAccount) Transfer(transferAmount float64, destinationAccount CheckingAccount) bool {
	if transferAmount < ca.balance && transferAmount > 0 {
		ca.balance -= transferAmount
		destinationAccount.Deposit(transferAmount)
		return true
	} else {
		return false
	}
}

func (ca *CheckingAccount) GetBalance() float64 {
	return ca.balance
}
