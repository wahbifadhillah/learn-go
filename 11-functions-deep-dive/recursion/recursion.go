package recursion

import "fmt"

func main() {

	fact := factorial(5)
	fmt.Println(fact)

}

func factorial(number int) int {
	// Solving with recursion
	if number == 0 {
		return 1
	}
	return number * factorial(number-1)
	// Solving factorial with loop
	// result := 1
	// for i := 1; i <= number; i++ {
	// 	result *= i
	// }

	// return result
}
