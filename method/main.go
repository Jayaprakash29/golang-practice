package main

import "fmt"

func main() {
	fmt.Println("Struct in golang")

	sathya := User{"sathya", "psathya@gmail.com", true, 05}
	fmt.Println(sathya)
	fmt.Printf("sathya in details v %+v\n", sathya)
	fmt.Printf("Name is %v, Email is %v", sathya.Name, sathya.Email)
	sathya.GetStatus()

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) GetStatus() {

	fmt.Println("is uer active: ", u.Status)
}
