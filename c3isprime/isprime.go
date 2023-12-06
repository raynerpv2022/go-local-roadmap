// package main Is a prime number?
package main

import "fmt"

// Isprime... return true only if argument number is prime
func Isprime(number int) bool {
	for i := 2; i <= number; i++ {
		if i != number {
			if number%i == 0 {
				return false
			}
		}

	}
	return true
}

// PrintPrimeNUmber show in a range  all prime number
func PrintPrimeNumber(ran int) {
	for i := 2; i <= ran; i++ {
		if Isprime(i) {
			fmt.Printf("%v ", i)
		}

	}
	fmt.Println()
}
func main() {
	PrintPrimeNumber(100)

}
