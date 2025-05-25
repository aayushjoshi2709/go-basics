package main

import (
	"fmt"
	"os"
	"strconv"
	"errors"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"

func getBalanceFromFile() (float64, error) {
	data, err := os.ReadFile(ACCOUNT_BALANCE_FILE)

	if err != nil {
		return 1000, errors.New("failed to read file")
	}

	balanceText := string(data)
	balance, err := strconv.ParseFloat(balanceText, 64)

	if err != nil {
		return 1000, errors.New("failed to parse balance file")
	}

	return balance, nil
}

func writeBalanceToAFile(balance float64) {
	balanceText := fmt.Sprint(balance)
	os.WriteFile(ACCOUNT_BALANCE_FILE, []byte(balanceText), 0644)
}

func main() {
	accountBalance, err  := getBalanceFromFile()

	if err != nil{
		fmt.Println("Error")
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Welcome to the GO Bank!")
	for {
		fmt.Println("What do you want to do?")
		fmt.Println("1. Check Balance")
		fmt.Println("2. Deposit Money")
		fmt.Println("3. Withdraw Money")
		fmt.Println("4. Exit")
		fmt.Print("Please enter your choice (1-4): ")
		var choice int
		fmt.Scan(&choice)

		fmt.Println("You chose option:", choice)

		switch choice {
		case 1:
			fmt.Printf("Your current balance is: %.2f\n", accountBalance)
		case 2:
			fmt.Print("Enter amount to deposit: ")
			var depositAmount float64
			fmt.Scan(&depositAmount)
			accountBalance += depositAmount
			writeBalanceToAFile(accountBalance)
			fmt.Printf("Balance updated successfully! Your new balance is: %.2f\n", accountBalance)
		case 3:
			fmt.Print("Enter amount to withdraw:")
			var withdrawlAmount float64
			fmt.Scan(&withdrawlAmount)
			if withdrawlAmount < 0 {
				fmt.Println("Invalid amount! Please enter a positive number.")
				continue
			}
			if withdrawlAmount <= accountBalance {
				accountBalance -= withdrawlAmount
				writeBalanceToAFile(accountBalance)
				fmt.Printf("Balance updated successfully! Your new balance is: %.2f\n", accountBalance)
			} else {
				fmt.Printf("insufficient funds! Your current balance is: %2f\n", accountBalance)
			}
		case 4:
			fmt.Println("Exiting the bank...")
			return
		default:
			fmt.Println("Invalid choice! Please select a valid option (1-4).")
		}
	}
}
