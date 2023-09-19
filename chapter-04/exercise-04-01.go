package main

import (
	"fmt"
)

var x bool

func main() {
	fmt.Println(x) // zero value
	x = true
	fmt.Println(x) // valor atribuido
	x := 10 < 100
	y := 10 == 100
	z := 10 > 100
	fmt.Println(x) // zero value
	fmt.Println(y) // zero value
	fmt.Println(z) // zero value
}
