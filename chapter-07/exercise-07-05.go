package main

import (
	"fmt"
)

func main() {

	for x := 10; x <= 100; x++ {
		fmt.Printf("number: %v, result divide per 4: %v \n", x, x%4)
	}
}
