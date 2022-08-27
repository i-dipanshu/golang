package main

import "fmt"

func main() {
	fmt.Println("Welcome to Maps")

	language := make(map[string]string)

	language["JS"] = "Javascript"
	language["Rb"] = "RUBY"
	language["java"] = "JAVA"

	fmt.Println("Language: ", language)
	fmt.Println("Short for java: ", language["JS"])

	delete(language, "RB")
	fmt.Println("Language: ", language)

}
