package accounts

import c "_projects/go-bank/clients"

type SavingsAccount struct {
	Client                   c.Client
	Agency                   string
	AccountNumber, Operation int
	balance                  float64
}

func (sa *SavingsAccount) Withdraw(amount float64) string {
	allowWithdrawal := amount > 0 && amount <= sa.balance
	if allowWithdrawal {
		sa.balance -= amount
		return "Withdrawal operation was successful."
	} else {
		return "Insuficient founds."
	}
}

func (sa *SavingsAccount) Debit(amount float64) string {
	allowWithdrawal := amount > 0 && amount <= sa.balance
	if allowWithdrawal {
		sa.balance -= amount
		return "Debit operation was successful."
	} else {
		return "Insuficient founds."
	}
}

func (sa *SavingsAccount) Deposit(amount float64) (string, float64) {
	if amount >= 0 {
		sa.balance += amount
		return "Deposit operation was successful.", sa.balance
	} else {
		return "Invalid value", sa.balance
	}
}

func (sa *SavingsAccount) GetBalance() float64 {
	return sa.balance
}
