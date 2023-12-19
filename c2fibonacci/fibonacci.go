// package main calculates fibonacci using for loop, using recursive function and factorial
package c2fibonacci

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

// function receive min and max , return  serie fibonacci between  both number
func FiboInArray(min, max int) [][]int {
	//  using [][]int because map is unordered data, is more dificult order it
	fiboList := [][]int{}
	for i := min; i <= max; i++ {
		temp := []int{}
		temp = append(temp, i, fibo(i))
		fiboList = append(fiboList, temp)
	}

	return fiboList
}

func fibo(n int) int {
	if n == 0 {
		return 0
	}
	if n < 3 {
		return 1
	}

	return fibo(n-1) + fibo(n-2)
	// f(1 ) = f(1)
	// f(2) = f(1) + f(0) = 1 + 0 = 1
	// f(3) = f(2) + f(1)  = 1 + 1 =2
	// f(4) = f(3) + f(2) = 2 +1 = 3
	// f(5) = f(4) + f(3) = 3 + 2 = 5
	// f(6) = f(5)  + f(4) = 5 + 3 = 8
}

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

// func main() {
// 	fmt.Println(FiboFor(7))

// 	fmt.Println(FiboRecur(7))
// 	fmt.Println(Facto(7))
// }
