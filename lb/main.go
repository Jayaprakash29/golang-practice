package main

import "fmt"

func main() {

	days := []string{"sunday", "monday", "tuesday", "Thursday", "friday", "saturday"}
	fmt.Println(days)

	// for i := 0; i < len(days); i++ {
	// 	fmt.Println(days)

	// }

	// for i := range days {
	// 	fmt.Println(days(i))
	//}
	for index, value := range days {
		fmt.Println(" index %v value %v\n ", index, value)

	}

}
