package main

import "fmt"

func largest(n ...int) int {
	var largest int

	for _, x := range n {
		if x > largest {
			largest = x
		}
	}

	return largest
}

func main() {
	biggest := largest(4, 21, 5, 63, 92, 71, 38)
	fmt.Println(biggest)
}
