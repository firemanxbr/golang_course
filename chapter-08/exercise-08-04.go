package main

import (
	"fmt"
)

func main() {

	//			    1.           2.           3.         4.                5.
	sabores := []string{"pepperoni", "mozzarela", "abacaxi", "quatroqueijos", "marg"}

	fatia := sabores[2:4]

	fmt.Println(fatia)
}
