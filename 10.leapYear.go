package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter the year : ")
	fmt.Scanln(&n)

	if n%4 == 0 {
		fmt.Println(n, "is a leap year ")
	} else {
		fmt.Println(n, "is not a leap year ")
	}
}
