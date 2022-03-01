package main

import "fmt"

func main() {
	x := 200
	y := x << 1

	fmt.Printf("%d\t %b\t %#x\n", x, x, x)
	fmt.Printf("%d\t %b\t %#x\n", y, y, y)
}
