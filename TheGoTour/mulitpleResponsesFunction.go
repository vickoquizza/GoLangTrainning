package main

import (
	"fmt"
)

func addAndRest(x, y int) (int, int) { // or func add(x,y int) int because x and y share the same type
	return x + y, x - y
}

func main() {
	a, b := addAndRest(3, 7)
	fmt.Printf("Adding numbers with a funcion %v, and making a rest at the same time %v", a, b)
}
