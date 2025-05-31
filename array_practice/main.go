package main

import "fmt"

type Product struct {
	title string
	id    int
	price int
}

func main() {
	hobbies_array := [3]string{"Programming", "Reading", "Sleeping"}
	fmt.Println(hobbies_array)
	fmt.Println(hobbies_array[0])
	fmt.Println(hobbies_array[1:])
	hobbies_slice1 := hobbies_array[:2]
	hobbies_slice2 := hobbies_array[0:2]
	fmt.Println(hobbies_slice1, hobbies_slice2)
	resliced_slice := hobbies_slice1[1:3]
	fmt.Println(resliced_slice)

	goals_arr := []string{"Sde2", "Sde3"}
	fmt.Println(goals_arr)
	goals_arr[1] = "sdm"
	fmt.Println(goals_arr)
	goals_arr = append(goals_arr, "vp")
	fmt.Println(goals_arr)

	product_arr := []Product{
		Product{
			title: "bag",
			id:    1,
			price: 100,
		},
		Product{
			title: "purse",
			id:    2,
			price: 220,
		},
	}
	fmt.Println(product_arr)

	// or
	product_arr = []Product{
		{
			title: "bag",
			id:    1,
			price: 100,
		},
		{
			title: "purse",
			id:    2,
			price: 220,
		},
	}
	fmt.Println(product_arr)

	// or
	product_arr = []Product{
		{"bag", 1, 100},
		{"purse", 2, 220},
	}
	fmt.Println(product_arr)

	product_arr = append(product_arr, Product{
		title: "chain",
		id:    3,
		price: 100,
	})
	fmt.Println(product_arr)
}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
