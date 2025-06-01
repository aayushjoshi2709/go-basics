package main

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	doubled := doubleNumbers(&numbers)
	fmt.Println(doubled)
	doubled = doubleNumbers2(&numbers)
	fmt.Println(doubled)

	// passing function to function
	doubled = transformNumbers(&numbers, double)
	tripled := transformNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)

	// returning function from function
	transformerFn := getTransformerFunction()
	doubled = transformNumbers(&numbers, transformerFn)
	fmt.Println(doubled)

}

func doubleNumbers(numbers *[]int) []int {
	doubleNumbers := []int{}
	for _, value := range *numbers {
		doubleNumbers = append(doubleNumbers, value*2)
	}
	return doubleNumbers
}

func doubleNumbers2(numbers *[]int) []int {
	doubleNumbers := []int{}
	for _, value := range *numbers {
		doubleNumbers = append(doubleNumbers, double(value))
	}
	return doubleNumbers
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	result := []int{}
	for _, value := range *numbers {
		result = append(result, transform(value))
	}
	return result
}

func getTransformerFunction() func(int) int {
	return double
}
