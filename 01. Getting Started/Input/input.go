package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input the rating : ")

	// comma ok || comma err
	input, err := reader.ReadString('\n')
	fmt.Print("Your rating: " + input)
	fmt.Print("Error = ", err)

}
