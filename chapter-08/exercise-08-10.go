package main

import (
	"fmt"
)

func main() {
	qualquercoisa := map[int]string{
		123: "muito legal",
		98:  "menos legal um pouquinho",
		983: "essa eh massa",
		18:  "idade de ir pra festa",
	}

	fmt.Println(qualquercoisa)

	total := 0

	for key, _ := range qualquercoisa {
		total += key
	}

	fmt.Println(total)
}
