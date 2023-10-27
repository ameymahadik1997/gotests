package main

import "fmt"

type BankAccount struct {
	ownerName      string
	accountBalance float64
}

func structToSlice(ba BankAccount) []BankAccount {
	var BankAccountSlice []BankAccount
	BankAccountSlice = append(BankAccountSlice, ba)
	return BankAccountSlice
}

func (ba *BankAccount) deposit(amount float64) {
	fmt.Print("\n------Amount: ", amount, " is deposited in the Account:------\n")
	if amount > 0 {
		ba.accountBalance += amount
	}
}

func (ba *BankAccount) balance() {
	BankAccountSlice := structToSlice(*ba)
	for _, banker := range BankAccountSlice {
		fmt.Print("Account Holder Name: ", banker.ownerName, " \nAccount Balance: ", banker.accountBalance, " Rs /-\n")
	}
}

func (ba *BankAccount) withdraw(amount float64) {
	fmt.Print("\n------Amount: ", amount, " is asked for withdrawal in the Account:------\n")
	if amount <= ba.accountBalance {
		ba.accountBalance -= amount
	} else {
		fmt.Print("\nWithdrawal Limit Reached\nAmount asked for withdrawal: ", amount, " is bigger than the amount in the Account: ", ba.accountBalance, "\n")
	}
}
