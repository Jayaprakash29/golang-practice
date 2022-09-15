package main

import "fmt"

func main() {

	var languages = make(map[string]string)

	languages["js"] = "javas"
	languages["PY"] = "python"
	languages["rb"] = "ruby"
	languages["HT"] = "html"

	fmt.Println("all programing languages are :", languages)
	fmt.Println("PY short is :", languages["PY"])

	for _, value := range languages {
		fmt.Printf("For key v, value is %v\n", value)

	}
}
