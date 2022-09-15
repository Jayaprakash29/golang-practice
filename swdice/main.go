package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	rand.Seed(time.Now().UnixNano())
	fmt.Println("Switch case in golang")

	dicenumber := rand.Intn(6) + 1
	fmt.Println("The number is ", dicenumber)

	switch dicenumber {

	case 1:
		fmt.Println("Boo! Its 1 You can start")
	case 2:
		fmt.Println("moe two step")
	case 3:
		fmt.Println("move 3 step")
		fallthrough
	case 4:
		fmt.Println("move 4 step")
		fallthrough
	case 5:
		fmt.Println("move 5 step")
	case 6:
		fmt.Println("BOOOOOM! Dice again")

	}

}
