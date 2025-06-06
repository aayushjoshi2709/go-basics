package main

import "fmt"

func main() {
	numbers := []int{1, 2, 3}
	transformed := transformedNumbers(&numbers, func(number int) int { return number * 2 })
	fmt.Println(transformed)

	double := createTransformer(2)
	triple := createTransformer(3)
	doubled := transformedNumbers(&numbers, double)
	tripled := transformedNumbers(&numbers, triple)
	fmt.Println(doubled)
	fmt.Println(tripled)
}

func transformedNumbers(numbers *[]int, transform func(int) int) []int {
	dNumbers := []int{}
	for _, val := range *numbers {
		dNumbers = append(dNumbers, transform(val))
	}
	return dNumbers
}

func createTransformer(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
