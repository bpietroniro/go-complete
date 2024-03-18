package main

import "fmt"

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

type Product struct {
	title string
	id    int64
	price float64
}

func main() {
	hobbies := [3]string{"hiking", "cooking", "games"}
	fmt.Println(hobbies)
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:3])

	hobbiesSlice := hobbies[:2]
	fmt.Println(hobbiesSlice) // ["hiking", "cooking"]
	// hobbiesSlice = []string{hobbies[0], hobbies[1]}

	hobbiesSlice = hobbiesSlice[1:3]
	fmt.Println(hobbiesSlice) // ["cooking", "games"]

	goals := []string{"learn", "build"}
	goals[1] = "master"
	goals = append(goals, "serve")
	fmt.Println(goals)

	products := []Product{
		{title: "pen", id: 1, price: 3.99},
		{title: "notebook", id: 2, price: 8.99},
	}

	products = append(products, Product{title: "eraser", id: 3, price: 0.99})
	fmt.Println(products)
}
