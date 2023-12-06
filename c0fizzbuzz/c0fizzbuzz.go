// Function receive 3 argument, 22 range beetwen 0 and n2, n0 and n1 are number to figure out if number in range are div by n0, n1 or both
// improve function to pass indeterminate argument and do the same
// improve  assing default value to argument

package c0fizzbuzz

import "fmt"

func Fizz_Buzz_100(num ...int) {

	for index := 0; index <= num[2]; index++ {
		if index%num[0] == 0 && index%num[1] == 0 {
			fmt.Println("FizzBuzz")

		} else if index%num[0] == 0 {
			fmt.Println("Fizz")
		} else if index%num[1] == 0 {
			fmt.Println("Buzz")
		} else {
			fmt.Println(index)
		}

	}

}
