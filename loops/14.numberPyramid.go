package main

import "fmt"

func main() {

	n := 10
	spaces := n - 1

	for i := 1; i <= n; i++ {
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
