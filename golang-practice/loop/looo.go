package main

import (
	"fmt"
)

func main() {
	// for i := 0; i <= 100; i += 10 {
	// 	fmt.Println(i)
	// }
	for i := 0; i < 5; i++ {
		if i == 3 {
			continue
		}
		fmt.Println(i)
	}

	//---------------nested loop------------------
	adj := [2]string{"big", "tasty"}
	fruits := [3]string{"apple", "orange", "banana"}
	for i := 0; i < len(adj); i++ {
		for j := 0; j < len(fruits); j++ {
			fmt.Println(adj[i], fruits[j])
		}
	}

}
