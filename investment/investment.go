package main

import (
	"fmt"
	"math"
)

// const always need to be initialized
const inflationRate = 2.5

func main() {

	var investmentAmmount, years float64 = 0, 10

	// this only works if you provide a value
	expectedReturnRate := 5.5

	fmt.Print("Enter the investment amount: ")
	fmt.Scan(&investmentAmmount)

	fmt.Print("Enter the number of years: ")
	fmt.Scan(&years)

	fmt.Print("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue, futureRealValue := calculateFutureValues(investmentAmmount, expectedReturnRate, years)
	printValues(futureValue, futureRealValue)

	futureValue, futureRealValue = otherCalculateFutureValues(investmentAmmount, expectedReturnRate, years)
	printValues(futureValue, futureRealValue)
}

func calculateFutureValues(investmentAmmount, expectedReturnRate, years float64) (float64, float64) {
	fv := investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
	rfv := fv / math.Pow(1+inflationRate/100, years)
	return fv, rfv
}

func otherCalculateFutureValues(investmentAmmount, expectedReturnRate, years float64) (fv float64, rfv float64) {
	fv = investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
	rfv = fv / math.Pow(1+inflationRate/100, years)
	return
}

func printValues(futureValue, futureRealValue float64){
	fmt.Printf("Future Value of Investment: $%.2f\n", futureValue)
	fmt.Printf("Future Real Value of Investment (adjusted for inflation): $%.2f\n", futureRealValue)
}