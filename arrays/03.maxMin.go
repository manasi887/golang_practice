package main

import "fmt"

func main() {

	var arr [5]int

	fmt.Println("Enter the element: ")

	for i := 0; i < 5; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Your array is: ", arr)

	max := arr[0]
	min := arr[0]

	for i := 1; i < 5; i++ {
		if max < arr[i] {
			max = arr[i]
		}
		if min > arr[i] {
			min = arr[i]
		}
	}

	fmt.Println("Minimum element: ", min)
	fmt.Println("Maximum element: ", max)
}
