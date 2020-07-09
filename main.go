package main

import "fmt"

func main() {
	nums := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	for _, even := range nums {
		if even%2 == 0 {
			fmt.Printf("%v is even\n", even)
		} else {
			fmt.Printf("%v is odd\n", even)
		}
	}
}
