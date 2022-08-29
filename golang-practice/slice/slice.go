package main

import "fmt"

func main() {

	// myslice1 := []int{}
	// fmt.Println(len(myslice1))
	// fmt.Println(cap(myslice1))
	// fmt.Println(myslice1)

	// myslice2 := []string{"Go", "Slices", "Are", "Powerful"}
	// fmt.Println(len(myslice2))
	// fmt.Println(cap(myslice2))
	// fmt.Println(myslice2)

	//This example shows how to create a slice from an array:---------------------------------

	// arr1 := [6]int{10, 11, 12, 13, 14, 15}
	// myslice := arr1[2:4]

	// fmt.Printf("myslice = %v\n", myslice)
	// fmt.Printf("length = %d\n", len(myslice))
	// fmt.Printf("capacity = %d\n", cap(myslice))
	//------------------------------------------------------

	//Create a Slice With The make() Function-----------------------------

	// myslice1 := make([]int, 5, 10)
	// fmt.Printf("myslice1 = %v\n", myslice1)
	// fmt.Printf("length = %d\n", len(myslice1))
	// fmt.Printf("capacity = %d\n", cap(myslice1))

	// // with omitted capacity
	// myslice2 := make([]int, 5)
	// fmt.Printf("myslice2 = %v\n", myslice2)
	// fmt.Printf("length = %d\n", len(myslice2))
	// fmt.Printf("capacity = %d\n", cap(myslice2))
	//------------------------------------------------------------------

	//append elements to the end of a slice------------
	// myslice3 := []int{1, 2, 3, 4, 5, 6}
	// fmt.Printf("myslice1 = %v\n", myslice3)
	// fmt.Printf("length = %d\n", len(myslice3))
	// fmt.Printf("capacity = %d\n", cap(myslice3))

	// myslice3 = append(myslice3, 20, 21)
	// fmt.Printf("myslice1 = %v\n", myslice3)
	// fmt.Printf("length = %d\n", len(myslice3))
	// fmt.Printf("capacity = %d\n", cap(myslice3))
	//----------------------------------------------------------

	arr1 := [6]int{9, 10, 11, 12, 13, 14} // An array
	myslice1 := arr1[1:5]                 // Slice array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = arr1[1:3] // Change length by re-slicing the array
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))

	myslice1 = append(myslice1, 20, 21, 22, 23) // Change length by appending items
	fmt.Printf("myslice1 = %v\n", myslice1)
	fmt.Printf("length = %d\n", len(myslice1))
	fmt.Printf("capacity = %d\n", cap(myslice1))
}
