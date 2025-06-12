package main

import "fmt"

func main() {
	var a, b, c, d, e float32

	fmt.Print("Enter the marks for 5 subjects : ")
	fmt.Scanln(&a)
	fmt.Scanln(&b)
	fmt.Scanln(&c)
	fmt.Scanln(&d)
	fmt.Scanln(&e)

	marks := a + b + c + d + e

	avg := marks / 5

	fmt.Println("The total marks are ", marks)

	if avg >= 90 {
		fmt.Println("Grade: O")
	} else if avg >= 80 {
		fmt.Println("Grade: A")
	} else if avg >= 70 {
		fmt.Println("Grade: B")
	} else if avg >= 60 {
		fmt.Println("Grade: C")
	} else if avg >= 50 {
		fmt.Println("Grade: D")
	} else if avg >= 40 {
		fmt.Println("Grade: E")
	} else {
		fmt.Println("Grade: F (Fail)")
	}
}
