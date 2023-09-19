package main

import (
	"fmt"
)

func main() {
	born := 1978
	current_year := 2023

	for {
		if born > current_year {
			break
		}

		fmt.Println(born)
		born++
	}
}
