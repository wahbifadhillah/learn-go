package main

import (
	"errors"
	"fmt"
	"os"
)

func main() {
	// var revenues, expenses, taxRate float64

	revenues, err := getInput("Revenue Amount: ")
	if err != nil {
		handleError(err)
		return
	}
	// fmt.Print("Revenue Amount: ")
	// fmt.Scan(&revenues)

	expenses, err := getInput("Expenses Amount: ")
	if err != nil {
		handleError(err)
		return
	}
	// fmt.Print("Expenses Amount: ")
	// fmt.Scan(&expenses)

	taxRate, err := getInput("Tax Rate %: ")
	if err != nil {
		handleError(err)
		return
	}
	// fmt.Print("Tax Rate %: ")
	// fmt.Scan(&taxRate)

	ebt, profit, ratio := calculateResults(revenues, expenses, taxRate)
	// earningBeforeTax := revenues - expenses
	// profit := earningBeforeTax * (1 - taxRate/100)
	// ratio := earningBeforeTax / profit

	fmt.Println("Earning Before Tax: ", ebt)
	fmt.Println("Profit: ", profit)
	fmt.Println("Ratio: ", ratio)
}

func getInput(text string) (variable float64, err error) {
	fmt.Print(text)
	fmt.Scan(&variable)
	if variable <= 0 {
		return 0, errors.New("Amount must greater than 0!")
	}
	return variable, nil
}

func calculateResults(revenues, expenses, taxRate float64) (ebt, profit, ratio float64) {
	ebt = revenues - expenses
	profit = ebt * (1 - taxRate/100)
	ratio = ebt / profit
	formattedResults := fmt.Sprintf("%v\n%v\n%v", ebt, profit, ratio)
	os.WriteFile("results.txt", []byte(formattedResults), 0644)

	return ebt, profit, ratio
}

func handleError(err error) {
	fmt.Println("ERROR")
	fmt.Println(err)
	fmt.Println("------")
}
