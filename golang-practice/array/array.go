package main

import (
	"fmt"
)

//This example declares two arrays (arr1 and arr2) with defined lengths-------------
// func main() {
// 	var arr1 = [3]int{1, 2, 3}
// 	arr2 := [5]int{4, 5, 6, 7, 8}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
//-------------------------------------------------------------

//This example declares two arrays (arr1 and arr2) with inferred lengths

// func main() {
// 	var arr1 = [...]int{1, 2, 3}
// 	arr2 := [...]int{4, 5, 6, 7, 8}

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
//--------------------------------------------------------------------------

//declares an array of strings:-------------------------------------------

// func main() {
// 	var cars = [4]string{"Volvo", "BMW", "Ford", "Mazda"}
// 	fmt.Print(cars)
//---------------------------------------------------------

//his example shows how to access the first and third elements in the prices array
// func main() {
// 	prices := [3]int{10,20,30}

// 	fmt.Println(prices[0])
// 	fmt.Println(prices[2])
//-----------------------------------------------

//This example shows how to change the value of the third element in the prices array:---------------------

// func main() {
// 	prices := [3]int{10,20,30}

// 	prices[2] = 50
// 	fmt.Println(prices)
//--------------------------------------------------

//Array Initialization----------------------
// func main() {
// 	arr1 := [5]int{} //not initialized
// 	arr2 := [5]int{1,2} //partially initialized
// 	arr3 := [5]int{1,2,3,4,5} //fully initialized

// 	fmt.Println(arr1)
// 	fmt.Println(arr2)
// 	fmt.Println(arr3)
//------------------------------------------------

//Initialize Only Specific Elements------------------
//   func main() {
// 	arr1 := [5]int{1:10,2:40}

// 	fmt.Println(arr1)
//------------------------------------------------

// Find the Length of an Array-------------------------------
func main() {
	arr1 := [4]string{"Volvo", "BMW", "Ford", "Mazda"}
	arr2 := [...]int{1, 2, 3, 4, 5, 6}

	fmt.Println(len(arr1))
	fmt.Println(len(arr2))
	//---------------------------------------------------------

}
