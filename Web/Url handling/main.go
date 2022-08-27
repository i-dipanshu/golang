package main

import (
	"fmt"
	"net/url"
)

const myurl = "https://google.com:3001/dogs?outh=true&paymentId=ewfddvd&lang=golang"

func main() {
	result, _ := url.Parse(myurl)
	fmt.Printf("Result is of type %T\n", result)

	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Port())

	for _, val := range result.Query(){
		fmt.Println(val)
	}

	newUrl := &url.URL{
		Scheme: "https",
		Host: "google.com",
		Path: "/routes",
	}

	fmt.Println(newUrl.String())
}
