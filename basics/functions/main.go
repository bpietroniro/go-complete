package main

import "fmt"

func main() {
	sum := sumup(1, 10, 15)
	fmt.Println(sum)

	numbers := []int{2, 7, 29}
	anotherSum := sumup(numbers...)
	fmt.Println(anotherSum)
}

func sumup(numbers ...int) int {
	sum := 0

	for _, val := range numbers {
		sum += val
	}

	return sum
}
