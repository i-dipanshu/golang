package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter your rating: ")
	input := bufio.NewReader(os.Stdin)

	rating, _ := input.ReadString('\n')

	fmt.Print("Your earlier rating: ", rating)

	// we're trim bc otherwise include a
	// new line which cant be converted to int
	numrating, err := strconv.ParseFloat(strings.TrimSpace(rating), 64)

	if err != nil {
		fmt.Print(err)
	} else {
		fmt.Print("Your new rating: ", numrating+1)
	}

}
