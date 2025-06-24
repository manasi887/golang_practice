package main

import "fmt"

func main() {
	var a, b int
	a = 34
	b = 21

	fmt.Printf("%d \n%d\n", a, b)
	for i := 0; i <= 7; i++ {
		c := a - b
		fmt.Println(c)

		a = b
		b = c
	}
}
