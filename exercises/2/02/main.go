package main

import "fmt"

func main() {

	half := func(n int) (int, bool) {
		var isEven bool
		x := n / 2
		y := n % 2

		if y == 0 {
			isEven = true
		}

		return x, isEven
	}

	fmt.Println(half(1))
	fmt.Println(half(2))
}
