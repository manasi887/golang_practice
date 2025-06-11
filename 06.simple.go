package main

import "fmt"

func main() {
	var p, r, t int

	fmt.Print("Enter the principal amount : ")
	fmt.Scanln(&p)

	fmt.Print("Enter the interest rate : ")
	fmt.Scanln(&r)

	fmt.Print("Enter the time period in years. : ")
	fmt.Scanln(&t)

	s := (p * r * t) / 100

	fmt.Println("the simple interest earned is : ", s)
}
