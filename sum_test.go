package sums

import "fmt"

func ExampleSum() {
	fmt.Println(Sum(0, 5))
	fmt.Println(Sum(0, 10))
	fmt.Println(Sum(2, 5))
	fmt.Println(Sum(2, 10))
	fmt.Println(Sum(2, 0))
	fmt.Println(Sum(2, -3))

	// Output:
	// 15
	// 55
	// 14
	// 54
	// 0
	// 0
}
