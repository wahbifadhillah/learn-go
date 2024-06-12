package main

import "fmt"

func main() {
	age := 32
	var agePointer *int = &age
	fmt.Println("Age:", *agePointer)
	editAgeToAdultYears((agePointer))
	fmt.Println(age)
}

func editAgeToAdultYears(age *int) {
	// return *age - 18
	*age -= 18
}
