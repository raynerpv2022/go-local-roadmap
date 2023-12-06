package main

import (
	"fmt"

	"github.com/raynerpv2022/retos/binarysearch"
	"github.com/raynerpv2022/retos/c0fizzbuzz"
	"github.com/raynerpv2022/retos/c1anagram"
	"github.com/raynerpv2022/retos/c8decimaltobinary"
)

func main() {
	fmt.Println(" Challenge 0 Fizz Buzz")
	c0fizzbuzz.Fizz_Buzz_100(2, 10, 200) //3 argument are need

	fmt.Println("****************************")

	fmt.Println(" Binary Search")
	binarysearch.BinarySearch(6)

	fmt.Println("****************************")

	c1anagram.Anagram("caso", "saco")
	c1anagram.Palindrom("arroz", "zorra")

	fmt.Println("****************************")
	fmt.Println("Decimal to binary")
	fmt.Printf("12 to binary %v \n", c8decimaltobinary.DecimalToBinaryByDiv(12))
	fmt.Printf("12 to binary %v \n", c8decimaltobinary.DecimalToBinaryByPow(12))

}
