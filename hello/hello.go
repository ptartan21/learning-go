package main

import (
	"fmt"

	"github.com/greetings"
)

func main() {
	message := greetings.Hello("Patrick")
	fmt.Println(message)
}
