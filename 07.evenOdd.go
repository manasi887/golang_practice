package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter the number : ")
	fmt.Scanln(&n)

	if n%2 == 0 {
		fmt.Println(n, "is an even number")
	} else {
		fmt.Println(n, "is a odd number")
	}
}
