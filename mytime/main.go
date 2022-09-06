package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("HI! MY TIME IS")
	presentTime := time.Now()
	fmt.Println(presentTime.Format("01-02-2006 15:04:05 tuesday"))
}
