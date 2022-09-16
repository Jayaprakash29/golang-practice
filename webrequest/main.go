package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const Url = "https://loc.dev"

func main() {
	fmt.Println("Web request")

	response, err := http.Get(Url)

	if err != nil {
		panic(err)
	}
	fmt.Printf("The rsponse is:  %T\n ", response)

	response.Body.Close()

	ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}
	content := string("databytes")

	fmt.Println(content)
	response.Body.Close()
}
