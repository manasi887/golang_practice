package main

import "fmt"

func main() {
	var a, b int
	a = 0
	b = 1

	fmt.Printf("%d \n%d\n", a, b)
	for i := 0; i <= 7; i++ {
		c := a + b
		fmt.Println(c)

		a = b
		b = c
	}
}
