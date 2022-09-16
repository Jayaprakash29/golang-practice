package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("TO my files")
	content := " This is the go file www.go.dev"

	file, err := os.Create("./mylogofile.txt")
	if err != nil {
		panic(err)
	}

	length, err := io.WriteString(file, content)
	if err != nil {
		panic(err)
	}
	fmt.Println("the length is: ", length)
	file.Close()
}
