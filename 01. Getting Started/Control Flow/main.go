package main

import "fmt"

func main() {
	fmt.Println("Welcome the Control Flow class")

	var isLogged bool = true

	if isLogged {
		fmt.Println("You can continue")
	} else {
		fmt.Println("You can't continue")
	}

	if num := 9; num < 9{
		fmt.Println("Continue")
	}else if num > 9{
		fmt.Println("Return")
	}else{
		fmt.Println("VIP Entry")
	}
}
