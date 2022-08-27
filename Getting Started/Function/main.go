package main

import "fmt"

func main() {
	fmt.Println(add(1, 2))
	fmt.Println(mul(1, 2, 3, 4, 5, 6))
}

func add(a int, b int) (int, string) {
	return (a + b), "The sum"
}

func mul(v...int) int {
	mul := 1
	for _, v := range v {
		mul *= v
	}
	return mul
}
