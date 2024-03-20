package functions_as_values

import "fmt"

type transformFn func(int) int

func main() {
	numbers := []int{1, 2, 3, 4}
	moreNumbers := []int{5, 2, 7}

	double := createTransformer(2)
	triple := createTransformer(3)

	doubled := transformNumbers(&numbers, double)
	fmt.Println(doubled)

	tripled := transformNumbers(&numbers, triple)
	fmt.Println(tripled)

	quadrupled := transformNumbers(&numbers, func(number int) int {
		return number * 4
	})
	fmt.Println(quadrupled)

	transformer1 := getTransformer(&numbers)
	transformer2 := getTransformer(&moreNumbers)

	fmt.Println(transformNumbers(&numbers, transformer1))
	fmt.Println(transformNumbers(&moreNumbers, transformer2))
}

func transformNumbers(numbers *[]int, transform transformFn) []int {
	transformed := []int{}

	for _, val := range *numbers {
		transformed = append(transformed, transform(val))
	}

	return transformed
}

func getTransformer(numbers *[]int) transformFn {
	if (*numbers)[0] == 1 {
		return double
	} else {
		return triple
	}
}

func double(number int) int {
	return number * 2
}

func triple(number int) int {
	return number * 3
}

// factory function
func createTransformer(factor int) transformFn {
	return func(number int) int {
		return number * factor
	}
}
