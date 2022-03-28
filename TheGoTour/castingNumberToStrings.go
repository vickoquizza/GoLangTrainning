package main

import (
	"fmt"
	"strconv"
)

var (
	number int = 42
)

func main() {
	var numberFloaty string
	numberFloaty = strconv.Itoa(number) //it converts an integer to a plain ASCII chararcter
	fmt.Printf("%v, %T", numberFloaty, numberFloaty)
}
