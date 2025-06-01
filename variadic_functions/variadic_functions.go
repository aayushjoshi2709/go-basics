package main

import (
	"fmt"
)

func main() {
	numbers := []int{2, 3, 5}
	sum := Sumup(numbers)
	fmt.Println(sum)

	sum2 := Sumup2(22, 33, 55, 66)
	fmt.Println(sum2)
	anotherSum := Sumup2(numbers[0], numbers[1:]...)
	fmt.Println(anotherSum)
}

func Sumup(numbers []int) int {
	sum := 0

	for _, number := range numbers {
		sum += number
	}
	return sum
}

func Sumup2(startingValue int, numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum + startingValue
}
