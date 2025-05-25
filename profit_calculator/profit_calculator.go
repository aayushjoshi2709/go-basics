package main

import (
	"fmt"
	"math"
)

func main() {
	var revenue, expenses, taxRate float64
	
	fmt.Print("Enter the total revenue: ")
	fmt.Scan(&revenue)
	fmt.Print("Enter the total expenses: ")
	fmt.Scan(&expenses)
	fmt.Print("Enter the tax rate: ")
	fmt.Scan(&taxRate)

	earningBeforeTax := revenue - expenses
	tax := earningBeforeTax * (taxRate / 100)
	earningAfterTax := earningBeforeTax - tax
	ratio := (earningBeforeTax / revenue) * 100

	fmt.Println("Earnings before tax: ", earningBeforeTax)
	fmt.Println("Total tax applied: ", tax)
	fmt.Println("Earnings after tax: ", earningAfterTax)
	fmt.Println("Earnings ratio: ", math.Round(ratio))

}
