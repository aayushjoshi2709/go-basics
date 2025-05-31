package main

import "fmt"

type FloatMap map[string]float64

func (m FloatMap) Output() {
	fmt.Println(m)
}

func main() {
	websites := []string{"https://google.com", "http://aws.com"}
	fmt.Println(websites)
	websites2 := map[string]string{
		"Google":              "https://google.com",
		"Amazon Web Services": "https://aws.com",
	}
	fmt.Println(websites2)
	fmt.Println(websites2["Google"])
	websites2["Linkedin"] = "https://linkedin.com"
	fmt.Println(websites2)
	delete(websites2, "Google")
	fmt.Println(websites2)

	courseRatings := make(map[string]float64, 10) // capacity of map
	courseRatings["go"] = 2.7
	courseRatings["react"] = 4.8
	fmt.Println(courseRatings)

	courseRatings2 := make(FloatMap, 10) // capacity of map
	courseRatings2["go"] = 2.7
	courseRatings2["react"] = 4.8
	courseRatings2.Output()

	for key, value := range courseRatings2 {
		fmt.Println(key, value)
	}
}
