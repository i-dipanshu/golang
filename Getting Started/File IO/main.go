package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	fmt.Println("Welcome to File I/O")
	write()
	read("./created.txt")
}

func write() {
	file, err := os.Create("./created.txt")
	checkNilErr(err)
	content := "Please write this in a new file"
	length, err := io.WriteString(file, content)
	fmt.Println("Length = ", length)
	defer file.Close()
}

func read(filepath string) {
	databyte, err := ioutil.ReadFile(filepath)
	checkNilErr(err)
	fmt.Println("Data in file = ", string(databyte))

}

func checkNilErr(err error) {
	if err != nil {
		panic(err)
	}
}
