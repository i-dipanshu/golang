package main

import "fmt"

func main() {
	fmt.Println("Welcome to Loops!")

	list := []string{"Java", "Javascript", "Ruby", "Python"}

	// Basic for loop
	for i:=0; i < len(list); i++{
		fmt.Println("List : ", list[i])
	}

	for index, lang := range list {
		fmt.Printf("At index %v We have %v\n", index, lang)
	}

	i:= 0;
	for  i > 3 {
		fmt.Println(list[i])
		i++
	}
}
