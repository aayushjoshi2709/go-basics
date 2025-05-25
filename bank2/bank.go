package main

import (
	"fmt"
	"aayushjoshi2709/bank2/fileops"
	"github.com/pallinder/go-randomdata"
)

const ACCOUNT_BALANCE_FILE = "balance.txt"

func main() {
	accountBalance, err := fileops.GetFloatFromFile(ACCOUNT_BALANCE_FILE, 0)

	if err != nil {
		fmt.Println("Error")
		fmt.Println(err)
		panic(err)
	}

	fmt.Println("Welcome to the GO Bank!")
	fmt.Println("Reach us 24/7 at: ", randomdata.PhoneNumber())
	for {
		presentOptions()
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
			fileops.WriteFloatToFile(ACCOUNT_BALANCE_FILE, accountBalance)
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
				fileops.WriteFloatToFile(ACCOUNT_BALANCE_FILE, accountBalance)
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
