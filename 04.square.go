package main

import "fmt"

func main() {
	var a int

	fmt.Println("Enter the number : ")
	fmt.Scan(&a)

	n := a * a

	fmt.Println("The square of ", a, "is", n)

	b := a * a * a

	fmt.Println("The cube of ", a, "is", b)
}
