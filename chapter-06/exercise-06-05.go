package main

import (
	"fmt"
)

func main() {
	whoami := "test"

	switch whoami {
	case "marcelo", "test":
		fmt.Println("marcelo")
		fallthrough
	case "nobody":
		fmt.Println("nobody")
	default:
		fmt.Println("help")
	}
}
