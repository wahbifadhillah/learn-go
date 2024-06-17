package main

import "fmt"

type Product struct {
	id    string
	title string
	price float64
}

func main() {

	hobbies := [3]string{"Code", "Skateboard", "Music"}
	// 1
	fmt.Println(hobbies)

	// 2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	// 3
	mainHobbies := hobbies[:2]
	fmt.Println(mainHobbies)

	// 4
	fmt.Println(cap(mainHobbies))
	mainHobbies = mainHobbies[1:3]
	fmt.Println(mainHobbies)

	// 5
	courseGoals := []string{"Learn GO", "REST API GO"}
	fmt.Println(courseGoals)

	// 6
	courseGoals[1] = "REST API GO For Work"
	courseGoals = append(courseGoals, "REST API GO")
	fmt.Println(courseGoals)

	// 7
	products := []Product{
		{
			"first-product",
			"First Product",
			5.99,
		},
		{
			"second-product",
			"Second Product",
			7.99,
		},
	}

	fmt.Println(products)

	// 8
	products = append(products, Product{"third-product", "Third Product", 8.99})
	fmt.Println(products)

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
