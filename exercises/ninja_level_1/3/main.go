package main

import (
	"fmt"
)

var x = 42
var y = "James Bond"
var z = true

func main() {

	s := fmt.Sprintf("My name is %v and I'm %v. It's %v, I have a license to kill", y, x, z)

	fmt.Println(s)
}
