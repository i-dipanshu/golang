package main

import "fmt"

func main() {
	fmt.Println("Welcome to the pointer class")

	// default value
	// var ptr *int
	// fmt.Print(ptr)	// --> <nil>

	// reference 
	num := 34
	ptr := &num

	// it prints the location
	fmt.Println(ptr)
	
	// it prints the value in the location
	fmt.Println(*ptr)

}
