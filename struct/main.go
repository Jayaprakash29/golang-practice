package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	sathya := User{"jaya", "psathya@gmail.com", true, 05}
	fmt.Println(sathya)
	fmt.Printf("sathya in details v %+v\n", sathya)
	fmt.Printf("Name is %v, Email is %v", sathya.Name, sathya.Email)
}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
