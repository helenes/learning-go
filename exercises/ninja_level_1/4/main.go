package main

import "fmt"

type newtype int

var x newtype

func main() {

	fmt.Println(x)
	fmt.Printf("%T\n", x)
	x = 42
	fmt.Println(x)
}
