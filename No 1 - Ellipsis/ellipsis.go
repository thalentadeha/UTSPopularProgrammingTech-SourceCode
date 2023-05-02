package main

import (
	"fmt"
)

// fungsi ellipsis
func sum(num ...int) int {

	total := 0

	for _, i := range num {
		total += i
	}

	return total

}

func main() {
	fmt.Println(sum(1, 2, 3, 4))
	fmt.Println(sum(1, 2, 3, 4, 5, 6, 7, 8, 9, 10))
}
