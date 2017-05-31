package main

import "fmt"

func main() {
	fmt.Println("Hello World")
	fmt.Println("Hello, my name is Hector Elenes")

	var name string
	fmt.Print("What is your name? Write it here: ")
	fmt.Scan(&name)
	fmt.Printf("Hello %v \n", name)

	var small, large int
	fmt.Print("Pick a large number: ")
	fmt.Scan(&large)
	fmt.Print("Pick a smaller number: ")
	fmt.Scan(&small)
	fmt.Println("Remainder is: ", large%small)

	for i := 0; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		}
	}

	total := 0
	for i := 1; i < 1000; i++ {
		if i%3 == 0 || i%5 == 0 {
			total = total + i
		}
	}
	fmt.Println(total)

}
