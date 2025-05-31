package main

import "fmt"

type product struct {
	title string
	id    string
	price float64
}

type TempratureData struct {
	day1 float64
	day2 float64
}

func main() {
	var productNames [4]string
	fmt.Println(productNames)
	productNames = [4]string{"A Book"}
	prices := [4]float64{10.99, 9.99, 45.99, 20.00}
	fmt.Println(prices)
	fmt.Println(productNames)

	productNames[3] = "A Carpet"
	fmt.Println(productNames)

	fmt.Println(prices[0])

	featuredPrices := prices[1:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[:3]
	fmt.Println(featuredPrices)

	featuredPrices = prices[2:]
	fmt.Println(featuredPrices)

	// negative index does not work in go
	// featuredPrices = prices[:-1]
	// fmt.Println(featuredPrices)
	highlightedPrices := featuredPrices[1:]
	fmt.Println(highlightedPrices)

	// slices are references of original array
	fmt.Println(prices)
	newSlice := prices[1:2]
	fmt.Println("newSlice: ", newSlice)
	prices[0] = 60
	fmt.Println(prices)
	fmt.Println(len(prices), cap(prices))
	fmt.Println(len(newSlice), cap(newSlice))

	// we can change the slice size to select more elements
	newSlice = newSlice[0:3]
	fmt.Println("newSlice: ", newSlice)

	// slices for dynamic arrays

	newPrices := []float64{10.99, 8.99}
	fmt.Println(newPrices[1])
	fmt.Println(newPrices[:1])
	updatedPrices := append(newPrices, 5.99, 33, 99.22)
	fmt.Println(updatedPrices)

	merged_prices := append(newPrices, newPrices...)
	fmt.Println(merged_prices)

	// create slice of size 2
	// make (slice, size, capacity) beyond capacity new array will be allocated
	userNames := make([]string, 2, 5)
	fmt.Println(userNames)
	userNames = append(userNames, "aayush")
	fmt.Println(userNames)
	userNames[0] = "mamta"
	userNames[1] = "neeraj"
	fmt.Println(userNames)

	for index, value := range userNames {
		fmt.Println(index, value)
	}

}
