package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Kevin")
	fmt.Println(message)
}
