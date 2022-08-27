package main

import "fmt"

func main() {
	fmt.Println("Welcome to the Array class")

	// intialize and then declare
	var fruitlist [4]string
	fruitlist[0] = "apple"
	fruitlist[1] = "mango"
	fruitlist[3] = "banana"

	fmt.Println("List = ", fruitlist)
	fmt.Println("List = ", len(fruitlist))


	// intialize and declare 
	var list = [3]int{1, 2 ,3}
	fmt.Print(list)
}
