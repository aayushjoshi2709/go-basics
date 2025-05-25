package main

import "fmt"

func main() {
	age := 32

	var agePointer *int = &age

	fmt.Println("Age: ", *agePointer)
	fmt.Println("Age *:", agePointer)

	adultYears := getAdultYears(agePointer)
	fmt.Println("Adult years: ", adultYears)

	deductAgeByOne(agePointer)
	fmt.Println("Age after deduction: ", age)
}

func getAdultYears(age *int) int {
	return *age - 18
}

func deductAgeByOne(age *int) {
	*age = *age - 1
}
