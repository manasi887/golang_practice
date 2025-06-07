package main

import "fmt"

func main() {

	a := 10
	b := 15

	fmt.Println("Before Swap")
	fmt.Println("a = ", a, "b = ", b)

	temp := a
	a = b
	b = temp

	fmt.Println("After Swap")
	fmt.Println("a = ", a, "b = ", b)

}
