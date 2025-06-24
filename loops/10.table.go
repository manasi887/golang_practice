package main

import "fmt"

func main() {
	var a int

	fmt.Print("Enter the number : ")
	fmt.Scanln(&a)

	for i := 1; i <= 10; i++ {
		n := a * i

		fmt.Println(n)
	}
}
