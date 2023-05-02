package main

import (
	"fmt"
)

func main() {

	//declare array 1 dan 2
	arr1 := [3]string{}
	arr2 := [3]string{}

	//for loop input
	fmt.Print("Input array 1: ")
	for i := 0; i < len(arr1); i++ {
		fmt.Scan(&arr1[i])
	}

	fmt.Print("Input array 2: ")
	for i := 0; i < len(arr2); i++ {
		fmt.Scan(&arr2[i])
	}

	//logic check array
	isSame := true

	for i := 0; i < len(arr1); i++ {
		if arr1[i] != arr2[i] {
			fmt.Printf("Index array ke %d tidak sama", i)
			fmt.Println()

			isSame = false
		}
	}

	if isSame {
		fmt.Println("Semua index dalam array 1 dan 2 sama")
	}

}
