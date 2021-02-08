package main

import (
	"fmt"

	"example.com/lib"
)

func main() {
	// Get a greeting message and print it.
	message := lib.Hello("Bruce")
	fmt.Println(message)
}