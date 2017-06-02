package main

import "fmt"

func main() {

	var a uint32 = 1
	var b uint32 = 2
	var c uint32 = 3
	var total uint32 = 2

	for c < 4000000 {
		c = a + b
		fmt.Println(c)
		if c%2 == 0 {
			total += c
		}
		a = b
		b = c
	}
	fmt.Println(total)
}
