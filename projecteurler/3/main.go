package main

import "fmt"

func main() {
	fmt.Println(largestPrimeFactor(851475143))
}

func getFactors(n int) []int {
	var factors []int
	for i := 1; i <= n; i++ {
		if n%i == 0 {
			factors = append(factors, i)
		}
	}
	return factors
}

func isPrime(num int) bool {
	var isprime bool
	if len(getFactors(num)) == 2 {
		isprime = true
	} else {
		isprime = false
	}
	return isprime
}

func largestPrimeFactor(n int) int {
	var primeFactors []int
	fmt.Println("Getting factors..")
	factors := getFactors(n)

	fmt.Println("Getting prime factors..")
	for _, num := range factors {
		if isPrime(num) == true {
			primeFactors = append(primeFactors, num)
		}
	}

	return primeFactors[len(primeFactors)-1]
}
