// package main calculates fibonacci using for loop, using recursive function and factorial
package main

import "fmt"

// FiboFor fibonacci using for loop
func FiboFor(number int) []int {
	arrayFibo := []int{0, 1}
	for i := 0; i <= number; i++ {

		if i > 1 {
			arrayFibo = append(arrayFibo, arrayFibo[i-1]+arrayFibo[i-2])

			//fmt.Println(arrayFibo)
		}

	}
	return arrayFibo
}

// FiboRecur fibonacci using recursive function
func FiboRecur(number int) int {

	if number <= 1 {
		return number
	}

	return FiboRecur(number-2) + FiboRecur(number-1)
}

//  5     (4)+                                               (3)  r3+r2 = 5
//  4     (3)+                         (2) r2+r1=r3
//  3     (2)+            (1)  r1+r1 = r2
//	2     1+ 0  r1+r0 = r1
//  1
//  0

// Factor calculates factor for a number
func Facto(number int) int {
	if number == 0 {
		return 1
	}
	return number * Facto(number-1)
}

// 4  4*(3)      r24
// 3  3*(2)      r6
// 2  2*(1)      r2
// 1  1* (==0)   r1

func main() {
	fmt.Println(FiboFor(7))

	fmt.Println(FiboRecur(7))
	fmt.Println(Facto(7))
}
