package main

import "fmt"

type floatMap map[string]float64

func (m floatMap) output() {
	fmt.Println(m)
}

func main() {
	// (Array type, length, allocated space)
	userNames := make([]string, 2, 5)

	userNames[0] = "Julie"
	userNames[1] = "Andy"

	userNames = append(userNames, "Max")
	userNames = append(userNames, "Abby")
	fmt.Println(userNames)

	// (Map type, allocated space)
	courseRatings := make(floatMap, 4)

	courseRatings["go"] = 4.7
	courseRatings["react"] = 4.8
	courseRatings["angular"] = 4.8
	courseRatings["java"] = 4.8

	courseRatings.output()

	// Loop through array
	for index, value := range userNames {
		fmt.Println(index, value)
	}

	for key, value := range courseRatings {
		fmt.Println(key, value)
	}
}
