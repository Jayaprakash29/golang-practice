package main

import (
	"fmt"
	"sort"
)

func main() {

	var flist = []string{"apple", "banana", "payaya"}
	fmt.Println("type of fruit %T\n", flist)

	flist = append(flist, "mango", "lithci")
	fmt.Println(flist)
	flist = append(flist[1:4])
	fmt.Println(flist)

	highscs := make([]int, 4)
	highscs[0] = 42
	highscs[1] = 64
	highscs[2] = 55
	highscs[03] = 99
	fmt.Println(highscs)

	highscs = append(highscs, 63, 45, 65)
	fmt.Println(highscs)

	sort.Ints(highscs)
	fmt.Println(highscs)

	var courses = []string{"javaa", "python", "html", "web", "ruby"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println(courses)
}
