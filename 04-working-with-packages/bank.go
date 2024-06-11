package main

import (
	"fmt"

	"github.com/Pallinder/go-randomdata"
	"github.com/wahbifadhillah/learn-go/tree/main/04-working-with-packages/fileops"
)

const accountBalanceFile = "balance.txt"

func main() {
	var accountBalance, err = fileops.GetFloatFromFile(accountBalanceFile)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		fmt.Println("------")
		return
		// panic("Can't continue, sorry.")
	}
	fmt.Println("Welcome to Go Bank!")
	fmt.Println("Reach us 24/7", randomdata.PhoneNumber())
	for {
		presentOptions()
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
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
			fmt.Println("Balance updated! New amount: ", accountBalance)
		case 3:
			withdraw := withdraw(accountBalance)
			if withdraw <= 0 || withdraw > accountBalance {
				continue
			}
			accountBalance -= withdraw
			fileops.WriteFloatToFile(accountBalance, accountBalanceFile)
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
