package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Welcome to time package")

	presentTime := time.Now()
	fmt.Println(presentTime)

	year := presentTime.Format("02-01-2006 Monday")
	fmt.Println(year)

	createdDate := time.Date(2056, time.April, 12, 12, 12, 12, 12, time.UTC)
	fmt.Print(createdDate)
}
