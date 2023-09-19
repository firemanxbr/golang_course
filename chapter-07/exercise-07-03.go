package main

import (
	"fmt"
)

func main() {
	born := 1978
	current_year := 2023

	for born <= current_year {
		fmt.Println(born)
		born++
	}
}
