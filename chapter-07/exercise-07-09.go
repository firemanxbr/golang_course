package main

import (
	"fmt"
)

func main() {
	esporteFavorito := "fazer nada"

	switch esporteFavorito {
	case "futebol":
		fmt.Println("Gol")
	case "pescar":
		fmt.Println("fish")
	case "fazer nada":
		fmt.Println("Uhuuu")
	}
}
