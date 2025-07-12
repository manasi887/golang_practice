package main

import "fmt"

func main() {

	n := 5
	spaces := n - 1

	for i := 1; i <= n; i++ {
		for j := 0; j < spaces; j++ {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Printf("* ")
		}
		fmt.Println()
		spaces--
	}

	spaces = 1

	for i := n - 1; i >= 1; i-- {
		for j := spaces; j > 0; j-- {
			fmt.Print("  ")
		}
		for j := 1; j <= i; j++ {
			fmt.Printf("* ")
		}
		for j := i - 1; j >= 1; j-- {
			fmt.Printf("* ")
		}
		fmt.Println()
		spaces++
	}
}
