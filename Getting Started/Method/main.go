package main

import "fmt"

func main() {
	ob := User{"Ankur", "ankur@gamil.com"}
	assign(ob)
	fmt.Printf("Name = %v , Email = %v \n", ob.Name, ob.Email)
}

type User struct {
	Name  string
	Email string
}

func assign(ob User) {
	ob.Name = "Dipanshu"
	ob.Email = "dipanshu@gmail.com"
	fmt.Printf("Name = %v , Email = %v \n", ob.Name, ob.Email)
}
