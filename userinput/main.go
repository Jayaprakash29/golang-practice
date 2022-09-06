package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welocme to my trainee group"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Hi please rete my training:")

	input, err := reader.ReadString('\n')
	fmt.Println("Thanks for rating", input, err)
	fmt.Printf("type of rating:%T", input)

}
