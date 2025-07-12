package main

import "fmt"

func main() {

	n := 5
	spaces := n - 1

	for i := 1; i <= n; i++ {
		for j := 0; j < spaces; j++ {
			fmt.Print(" ")
		}
		for j := 1; j <= i; j++ {
			fmt.Print("* ")
		}
		fmt.Println()
		spaces--
	}
}
