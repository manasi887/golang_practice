package main

import "fmt"

func main() {
	n := 112
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
		fmt.Print("Is prime number")
	} else {
		fmt.Print("Not a prime number")
	}
}
