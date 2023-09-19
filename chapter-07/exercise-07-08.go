package main

import (
	"fmt"
)

func main() {
	x := 2

	switch {
	case x == 2:
		fmt.Println("x is", x)
	case x == 1:
		fmt.Println("x is", x)
	case x == 0:
		fmt.Println("x is", x)
	}
}
