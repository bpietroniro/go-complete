package recursion

import "fmt"

func main() {
	fact := factorial(6)
	fmt.Println(fact)

	fmt.Println(factorialRecursive(6, map[int]int{}))
}

func factorial(number int) int {
	result := 1

	for i := 1; i <= number; i++ {
		result *= i
	}

	return result
}

func factorialRecursive(number int, memo map[int]int) int {
	if number == 0 {
		return 1
	}

	if memo[number] != 0 {
		return memo[number]
	}

	memo[number] = number * factorialRecursive(number-1, memo)
	return memo[number]
}
