package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

const accountBalanceFile = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(accountBalanceFile)
	if err != nil {
		return 1000, errors.New("Failed to read balance file")
	}
	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)
	if err != nil {
		return 1000, errors.New("Failed to parse stored balance value")
	}
	return balance, nil
}

func writeBalanceToFile(balance float64) {
	formatBalance := fmt.Sprint(balance)
	os.WriteFile(accountBalanceFile, []byte(formatBalance), 0644)
}

func main() {
	var accountBalance, err = getBalanceFromFile()
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		return
		// panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")
	for {
		printMenu()
		var choice int
		fmt.Print("Your choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Println("Your balance is: ", accountBalance)
		case 2:
			deposit := deposit()
			if deposit <= 0 {
				continue
			}
			accountBalance += deposit
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		case 3:
			withdraw := withdraw(accountBalance)
			if withdraw <= 0 || withdraw > accountBalance {
				continue
			}
			accountBalance -= withdraw
			writeBalanceToFile(accountBalance)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		default:
			fmt.Println("Goodbye!")
			fmt.Println("Thanks for choosing our Bank!")
			return
			// break
		}

		// fmt.Println("Your choice is: ", choice)
	}
}

func printMenu() {
	fmt.Println("What do you want to do?")
	fmt.Println("1. Check balance")
	fmt.Println("2. Deposit money")
	fmt.Println("3. Withdraw money")
	fmt.Println("4. Exit")
}

func deposit() (depositAmount float64) {
	fmt.Print("Your deposit amount: ")
	fmt.Scan(&depositAmount)
	if depositAmount <= 0 {
		fmt.Println("Invalid amount! Must be greater than 0!")
		return -1
	}
	return depositAmount
}

func withdraw(accountBalance float64) (withdrawAmount float64) {
	fmt.Print("Your withdraw amount: ")
	fmt.Scan(&withdrawAmount)
	if withdrawAmount <= 0 {
		fmt.Println("Invalid amount! Must be greater than 0!")
		return -1
	} else if withdrawAmount > accountBalance {
		fmt.Println("You don't have succifient balance!")
		return -1
	}
	return withdrawAmount
}
