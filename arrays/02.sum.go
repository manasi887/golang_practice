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

	sum := 0

	for i := 0; i < 5; i++ {
		sum += arr[i]
	}

	avg := sum / 5

	fmt.Println("The sum is: ", sum)
	fmt.Println("The avg is: ", avg)
}
