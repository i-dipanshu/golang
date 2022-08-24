package main

import "fmt"

// we can't define a walrus operator outside a func

// Here capital P means public variable and const means it can't be changed
const Password = "2332HU";

func main() {

	// string
	var username string = "Dipanshu"
	fmt.Println("username")
	fmt.Printf("Variable is of type : %T", username)

	// boolean
	var isTrue bool = true
	fmt.Println(isTrue)
	fmt.Printf("Variable is of type : %T", isTrue)

	// implicit
	var num = 23
	fmt.Println(num)

	// walrus 
	walrus := "Hello"
	fmt.Println(walrus)

	// default values
	var name string
	fmt.Println(name)

	var nums int
	fmt.Println(nums)

}
