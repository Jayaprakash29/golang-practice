package main

import (
	"fmt"
	"strconv"
)

func main() {
	var arr = []int{1, 2, 3, 4, 5, 6}
	var n int = len(arr)
	var sum int = 7
	returnAllCombinations(arr, n, sum)

}

func returnAllCombinations(arr []int, n int, sum int) {
	// int count = 0;
	// Consider all possible pairs
	// and check their sums
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if arr[i]+arr[j] == sum {
				fmt.Println("(" + strconv.Itoa(i) + ", " + strconv.Itoa(j) + ")")
			}
		}
	}
}
