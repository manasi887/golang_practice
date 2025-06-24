package main

import "fmt"

func main() {
	var n int

	fmt.Print("Enter the number ")
	fmt.Scanln(&n)

	max := n / 2
	isPrime := true

	if n <= 1 {
		fmt.Println("Neither prime nor composite")
	} else {
		for i := 2; i <= max; i++ {
			if n%i == 0 {
				isPrime = false
				break
			}
		}

	}
	if isPrime {
		fmt.Println("Is prime number")
	} else {
		fmt.Println("Not a prime number")
	}
}
