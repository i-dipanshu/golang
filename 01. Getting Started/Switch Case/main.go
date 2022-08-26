package main

import "fmt"

func main() {
	fmt.Println("Welcome to switch case")

	switch num := 3; num {
	case 1:
		fmt.Println("Number is greater.")
	case 2:
		fmt.Println("Number is greater.")
	case 3:
		fmt.Println("Number Hitted, Falltrough.")
		fallthrough
	case 4:
		fmt.Println("Number is smaller.")
	case 5:
		fmt.Println("Number is smaller.")
	default:
		fmt.Println("Invalid input")

	}
}
