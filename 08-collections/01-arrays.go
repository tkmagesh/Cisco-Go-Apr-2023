package main

import "fmt"

func main() {
	// var nos [5]int // memory is allocated & initialized (with default value of int - 0)
	var nos [5]int = [5]int{3, 1, 4, 2, 5}
	// var nos [5]int = [5]int{3, 1, 4}

	// type inference
	// var nos = [5]int{3, 1, 4}
	// nos := [5]int{3, 1, 4}
	fmt.Println(nos)

	fmt.Println()
	fmt.Println("Iterating an array [using indexer]")
	for i := 0; i < 5; i++ {
		fmt.Printf("nos[%d] = %d\n", i, nos[i])
	}

	fmt.Println()
	fmt.Println("Iterating an array [using range]")
	for idx, val := range nos {
		fmt.Printf("nos[%d] = %d\n", idx, val)
	}

	sort(&nos)
	fmt.Println(nos)

	x := [5]int{3, 1, 4, 2, 5}
	y := [5]int{3, 1, 4, 2, 5}
	fmt.Printf("&x = %p\n", &x)
	fmt.Printf("&y = %p\n", &y)
	fmt.Println(x == y)

}

/* write sort fn to sort an array of 5 integers (bubble sort) */
func sort(x *[5]int) {
	for i := 0; i < 4; i++ {
		for j := i + 1; j < 5; j++ {
			if x[i] > x[j] {
				x[i], x[j] = x[j], x[i]
			}
		}
	}
}
