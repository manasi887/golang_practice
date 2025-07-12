package main

import "fmt"

func main() {

	n := 10
	spaces := 1
	max := n

	for i := 1; i <= n; i++ {
		for j := 1; j < spaces; j++ {
			fmt.Print("  ")
		}
		for j := 1; j <= max; j++ {
			fmt.Printf("%d ", j)
		}
		for j := max - 1; j >= 1; j-- {
			fmt.Printf("%d ", j)
		}
		fmt.Println()
		spaces++
		max--
	}
}

// package main

// import "fmt"

// func main() {

// 	n := 10
// 	spaces := 1

// 	for i := n; i >= 1; i-- {
// 		for j := spaces; j > 0; j-- {
// 			fmt.Print("  ")
// 		}
// 		for j := 1; j <= i; j++ {
// 			fmt.Printf("%d ", j)
// 		}
// 		for j := i - 1; j >= 1; j-- {
// 			fmt.Printf("%d ", j)
// 		}
// 		fmt.Println()
// 		spaces++
// 	}
// }
