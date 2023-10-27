package main

import "fmt"

func main() {
	fmt.Print("=================Report================\n")
	fmt.Print("The account program starts from here:\n")
	account := BankAccount{"Amay Mahadik", 1000.00}
	account.balance()
	account.deposit(2000.00)
	fmt.Print("\nThe updated balance:\n")
	account.balance()
	account.withdraw(900.00)
	fmt.Print("\nThe updated balance:\n")
	account.balance()
	account.withdraw(6900.00)
	fmt.Print("\nThe updated balance:\n")
	account.balance()
	fmt.Print("=======================================\n")
}
