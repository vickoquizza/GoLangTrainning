package main

import "fmt"

var (
	number int = 43
)

func main() {
	var numberFloaty float32
	numberFloaty = float32(number)
	fmt.Printf("%v, %T", numberFloaty, numberFloaty)
}
