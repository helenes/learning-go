package main

import "fmt"

const (
	atyped    int = 21
	anontyped     = "not typed"
)

func main() {
	number := 42
	fmt.Printf("%d\t%b\t%#X\n", number, number, number)
	numbershifted := number << 1
	fmt.Printf("%d\t%b\t%#X\n", numbershifted, numbershifted, numbershifted)

	fmt.Println(atyped)
	fmt.Println(anontyped)
}
