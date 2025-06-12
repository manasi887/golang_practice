package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter the number : ")
	fmt.Scanln(&n)

	if n > 0 {
		fmt.Println(n, "is a positive number")
	} else if n < 0 {
		fmt.Println(n, "is a negative number")
	} else {
		fmt.Println("Number is zero")
	}
}
