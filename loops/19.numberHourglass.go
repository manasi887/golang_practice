package main

import "fmt"

func main() {

	n := 5
	spaces := 1

	for i := n; i >= 1; i-- {
		for j := spaces; j > 0; j-- {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
		spaces++
	}

	spaces = n - 1

	for i := 2; i <= n; i++ {
		for j := 0; j < spaces; j++ {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("%d ", j)
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
		spaces--
	}
}
