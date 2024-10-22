package main

import (
	"fmt"

	greetings "github.com/brentgroves/greetings3"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Gladys")
	fmt.Println(message)
}
