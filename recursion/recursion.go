package main

import "fmt"

func main() {
	fact := factorial(5)
	fmt.Println(fact)

	fact = factorial2(5)
	fmt.Println(fact)
}

func factorial(number int) int {
	result := 1
	for i := 1; i <= number; i++ {
		result = result * i
	}
	return result
}

func factorial2(number int) int {
	if number == 1 || number == 0 {
		return number
	}

	return number * factorial2(number-1)
}
