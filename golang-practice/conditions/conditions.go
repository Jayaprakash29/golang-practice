package main

import (
	"fmt"
)

func main() {
	//------------------if-----------
	// 	x := 20
	// 	y := 18
	// 	if x > y {
	// 		fmt.Println("x is greater than y")
	// 	}
	//-----------------------------------------

	//----------------if else----------------------------
	// temperature := 14
	// if temperature > 15 {
	// 	fmt.Println("It is warm out there")
	// } else {
	// 	fmt.Println("It is cold out there")
	// }
	//-----------------------------------------------------

	//------------------else if------------------
	// a := 14
	// b := 14
	// if a < b {
	// 	fmt.Println("a is less than b.")
	// } else if a > b {
	// 	fmt.Println("a is more than b.")
	// } else {
	// 	fmt.Println("a and b are equal.")
	// }

	//-----------------------nested if------------------
	num := 20
	if num >= 10 {
		fmt.Println("Num is more than 10.")
		if num > 15 {
			fmt.Println("Num is also more than 15.")
		}
	} else {
		fmt.Println("Num is less than 10.")
	}
}
