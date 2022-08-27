package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url = "http://127.0.0.1:5500/Web%20Request/demo.html"

func main() {
	fmt.Println("Web Request demo")

	res, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()

	content, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(content))

	fmt.Println(res.StatusCode)
	fmt.Println(res.ContentLength)
	fmt.Println(res.Header)
	fmt.Println(res.Cookies())
}
