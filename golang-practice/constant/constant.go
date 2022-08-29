package main

import (
	"fmt"
)

//Declaring a Constant----------------------------
// const PI = 3.14
// func main() {
// 	fmt.Println(PI)
//---------------------------------

//Typed constants are declared with a defined type:-------------
// const A int = 1
// func main() {
//   fmt.Println(A)
//------------------------------------------------

//Untyped constants are declared without a type:----------------------
// const A = 1
// func main() {
//   fmt.Println(A)
//---------------------------------------------

// Multiple constants can be grouped together into a block for readability:
const (
	A int = 1
	B     = 3.14
	C     = "Hi!"
)

func main() {
	fmt.Println(A)
	fmt.Println(B)
	fmt.Println(C)
	//---------------------------------------------------------

}
