package main

import (
	"fmt"
)

func main() {
	amigos := map[string]int{
		"alfredo": 5551234,
		"joana":   9996674,
	}

	fmt.Println(amigos)
	fmt.Println(amigos["joana"])

	amigos["gopher"] = 4444444
	fmt.Println(amigos)
	fmt.Println(amigos["gopher"])

	if sera, ok := amigos["fantasma"]; !ok {
		fmt.Println("nao tem!")
	} else {
		fmt.Println(sera)
	}
}
