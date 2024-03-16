package main

import "fmt"

func main() {
	age := 32
	fmt.Println("Age:", age)

	changeYearsToAdultYears(&age)
	fmt.Println("Age overwritten:", age)
}

func changeYearsToAdultYears(age *int) {
	*age = *age - 18
}
