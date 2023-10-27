package main

import "fmt"

func main() {
	fmt.Print("=================Report================\n")
	fmt.Print("The account program starts from here:\n")
	account := BankAccount{"Amay Mahadik", 1000.00}
	account.print()
	account.deposit(2000.00)
	fmt.Print("\nThe updated balance:\n")
	account.print()
	account.withdraw(900.00)
	fmt.Print("\nThe updated balance:\n")
	account.print()
	account.withdraw(6900.00)
	fmt.Print("\nThe updated balance:\n")
	account.print()
	fmt.Print("=======================================\n")
}
