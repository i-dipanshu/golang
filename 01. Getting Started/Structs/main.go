package main

import "fmt"

func main() {
	s1 := User{"Ram", 14, 01}
	fmt.Println(s1)
	fmt.Printf("Student 1 : %+v\n",s1)
	fmt.Printf("Student 1 : Name = %v",s1.Name)
}

type User struct {
	Name    string
	Roll    int
	Section int
}
