package main

import "fmt"

func main(){
	result := add(22, 33)
	// error -> result += 1

	result2 := add2(22, 33)
	result2 += 1
	fmt.Println(result)
}

func add(a, b interface{}) interface{}{
	aInt, aIsInt := a.(int)
	bInt, bIsInt := b.(int)
	if aIsInt && bIsInt{
		return aInt + bInt
	}

	aFloat, aIsFloat := a.(float64)
	bFloat, bIsFlaat := b.(float64)
	if bIsFlaat && aIsFloat{
		return aFloat + bFloat
	}
	return nil
}


// we can also do -> func add2[T int|float64](a, b T) 
func add2[T int|float64](a, b T) T{
	return a + b
}