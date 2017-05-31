package main

import "fmt"

func half(n int) (int, bool) {

	var isEven bool
	x := n / 2
	y := n % 2

	if y == 0 {
		isEven = true
	}

	return x, isEven
}

func main() {

	fmt.Println(half(1))
	fmt.Println(half(2))
}
