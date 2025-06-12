package main

import "fmt"

func main() {
	var c float32

	fmt.Print("Enter the tempature is Celsius : -")
	fmt.Scan(&c)

	f := c*1.8 + 32

	fmt.Println("The tempature in Fahrenheit is : ", f)
}
