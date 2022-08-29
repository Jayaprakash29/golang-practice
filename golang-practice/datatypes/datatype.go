package main

import (
	"fmt"
)

//different data types in Go:--------------
// func main() {
// 	var a bool = true
// 	var b int = 5
// 	var c float32 = 3.14
// 	var d string = "Hi!"

// 	fmt.Println("Boolean: ", a)
// 	fmt.Println("Integer: ", b)
// 	fmt.Println("Float:   ", c)
// 	fmt.Println("String:  ", d)
//-------------------------------------------

// different ways to declare Boolean variables:
// func main() {
// 	var b1 bool = true
// 	var b2 = true
// 	var b3 bool
// 	b4 := true
// 	fmt.Println(b1)
// 	fmt.Println(b2)
// 	fmt.Println(b3)
// 	fmt.Println(b4)
//--------------------------------------------

// The string data type is used to store a sequence of characters
func main() {
	var txt1 string = "Hello!"
	var txt2 string
	txt3 := "World 1"

	fmt.Printf("Type: %T, value: %\n", txt1, txt1)
	fmt.Printf("Type: %T, value: %v\n", txt2, txt2)
	fmt.Printf("Type: %T, value: %v\n", txt3, txt3)

}
