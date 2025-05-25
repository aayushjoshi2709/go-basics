// package main

// import "fmt"
// import "math"

// func main(){
// 	var investmentAmmount float64 = 1000
// 	var expectedReturnRate = 5.5
// 	var years = 10
// 	var futureValue = investmentAmmount * math.Pow(1 + expectedReturnRate/100, float64(years))
// 	fmt.Println("Future Value of Investment: $", futureValue)
// }

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

	outputText("Enter the expected return rate: ")
	fmt.Scan(&expectedReturnRate)

	futureValue := investmentAmmount * math.Pow(1+expectedReturnRate/100, years)
	futureRealValue := futureValue / math.Pow(1+inflationRate/100, years)

	fmt.Println("Future Value of Investment: $", futureValue)
	fmt.Println("Future Real Value of Investment (adjusted for inflation): $", futureRealValue)

	fmt.Printf("Future Value of Investment: %.2f\n", futureValue)

	// sprint varients to send response to string

	formattedFV := fmt.Sprintf("Future Value of Investment: $%.2f\n", futureValue)
	formattedRFV := fmt.Sprintf("Future Real Value of Investment (adjusted for inflation): $%.2f\n", futureRealValue)
	fmt.Print(formattedRFV)
	fmt.Print(formattedFV)
}

func outputText(text string) {
	fmt.Print(text)
}
