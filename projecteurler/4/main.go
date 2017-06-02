// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.
// Find the largest palindrome made from the product of two 3-digit numbers.
package main

import "fmt"

func main() {
	fmt.Println("Project Euler's Problem 4")
	var largestResult int
	for i := 100; i <= 999; i++ {
		for j := 100; j <= 999; j++ {
			result := i * j
			if isPalindrome(result) {
				if result > largestResult {
					largestResult = result
				}
			}
		}
	}
	fmt.Println("Answer:", largestResult)

}

func isPalindrome(n int) bool {
	reversedn := reverseInt(n)
	palindrome := false
	if n == reversedn {
		palindrome = true
	}
	return palindrome
}

func reverseInt(n int) int {
	newint := 0
	for n > 0 {
		remainder := n % 10
		newint *= 10
		newint += remainder
		n /= 10
	}
	return newint
}
