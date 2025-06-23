package main

import "fmt"

func main() {
	n := 6
	sum := 0

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("The sum is ", sum)
}
