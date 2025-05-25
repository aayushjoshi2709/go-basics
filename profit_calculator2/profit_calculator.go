package main

import (
	"errors"
	"fmt"
	"math"
	"os"
)

func main() {
	revenue, err := InputAndReturn("Enter the total revenue: ")
	if err != nil {
		fmt.Println("Revenue must not be nill")
	}
	expenses, err := InputAndReturn("Enter the total expenses: ")
	if err != nil {
		fmt.Println("Revenue must not be nill")
	}
	taxRate, err := InputAndReturn("Enter the tax rate: ")
	if err != nil {
		fmt.Println("Revenue must not be nill")
	}

	earningBeforeTax, tax, earningAfterTax, ratio := calculateFinancials(revenue, expenses, taxRate)

	fmt.Println("Earnings before tax: ", earningBeforeTax)
	fmt.Println("Total tax applied: ", tax)
	fmt.Println("Earnings after tax: ", earningAfterTax)
	fmt.Println("Earnings ratio: ", math.Round(ratio))
}

func InputAndReturn(prompt string) (float64, error) {
	fmt.Print(prompt)
	var userInput float64
	fmt.Scan(&userInput)

	if userInput <= 0 {
		return 0, errors.New("Value must be a +ve number")
	}

	return userInput, nil
}

func calculateFinancials(revenue, expenses, taxRate float64) (float64, float64, float64, float64) {
	earningBeforeTax := revenue - expenses
	tax := earningBeforeTax * (taxRate / 100)
	earningAfterTax := earningBeforeTax - tax
	ratio := (earningBeforeTax / revenue) * 100
	writeToFile(earningBeforeTax, tax, earningAfterTax, ratio)
	return earningBeforeTax, tax, earningAfterTax, ratio
}

func writeToFile(earningBeforeTax, tax, earningAfterTax, ratio float64) {
	result := fmt.Sprintf("-----------------------------------\nEBT: %.2f\nTAX: %.2f\nEAT: %.2f\nRATIO: %.2f\n", earningBeforeTax, tax, earningAfterTax, ratio)
	os.WriteFile("Results.txt", []byte(result), 0644)
}
