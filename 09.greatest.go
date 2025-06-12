package main

import "fmt"

func main() {
	var a, b, c int

	fmt.Print("Enter three numbers : ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)

	if a > b && a > c {
		fmt.Println(a, "is the greatest number")
	} else if b > a && b > c {
		fmt.Println(b, "is the greatest number")
	} else if c > a && c > b {
		fmt.Println(c, "is the greatest number")
	} else {
		fmt.Println("All three numbers are equal")
	}
}
