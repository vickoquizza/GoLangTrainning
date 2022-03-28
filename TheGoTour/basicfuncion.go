package main

import (
	"fmt"
)

func add(x int, y int) int { // or func add(x,y int) int because x and y share the same type
	return x + y
}

func main() {
	fmt.Println("Adding numbers with a funcion", add(3, 7))
}
