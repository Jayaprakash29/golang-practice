package main

import "fmt"

func main() {
	// var ptr *int
	// fmt.Println("New to pointers", ptr)
	myNumber := 23
	var ptr = &myNumber
	fmt.Println("show pointer values", ptr)
	fmt.Println("show pointer values", *ptr)

	*ptr = *ptr / 2
	fmt.Println("print new values:", myNumber)

}
