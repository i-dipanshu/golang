package main

import "fmt"

func main() {
	displayReverse()
	defer fmt.Print("World ")
	fmt.Print("Hello ")
	defer fmt.Print("Golang ")
}

func displayReverse() {
	for i := 0; i < 3; i++ {
		fmt.Print(i, " --> ")
	}
}
