package main

import "fmt"

func main() {
	n := 6
	facotrial := 1

	for i := 1; i <= n; i++ {
		facotrial *= i
	}

	fmt.Println("The sum is ", facotrial)
}
