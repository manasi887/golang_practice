package main

import "fmt"

func main() {

	var arr [5]int

	fmt.Println("Enter the elements:")

	for i := 0; i < 5; i++ {
		fmt.Printf("Element %d: ", i+1)
		fmt.Scan(&arr[i])
	}

	fmt.Println("Your array is:", arr)
}
