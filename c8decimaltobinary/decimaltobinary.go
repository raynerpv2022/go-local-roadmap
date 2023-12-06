package c8decimaltobinary

import (
	"math"
)

// func main() {

// 	fmt.Println("Convert Decimal to Binary: ")
// 	for true {
// 		var input string
// 		fmt.Print("Enter number or press `quit` to quit...:")
// 		_, err := fmt.Scanln(&input)

// 		if err != nil {
// 			fmt.Println("Try again : ", err)
// 			continue
// 		}
// 		if input == "quit" {
// 			break
// 		}
// 		n, err := strconv.Atoi(input)
// 		if err != nil {
// 			fmt.Println("We have some error, Try again ")
// 			continue
// 		}
// 		fmt.Printf("%v : %v \n", n, decimalToBinary(n))
// 		fmt.Printf("%v : %v \n", n, decimalToBinary1(n))

// 	}

// }
func DecimalToBinaryByDiv(n int) []int {
	var resto int
	var binary []int
	for n > 0 {
		resto = n % 2
		n /= 2
		binary = append(binary, resto)

	}
	// reversing binary
	for i, j := 0, len(binary)-1; i < j; i, j = i+1, j-1 {
		binary[i], binary[j] = binary[j], binary[i]
	}

	return binary

}
func DecimalToBinaryByPow(n int) []int {

	exp := 0
	binary := []int{}
	sum := 0

	for n > 0 {
		if n < int(math.Pow(2, float64(exp+1))) {

			for ; exp >= 0; exp-- {
				if (sum + int(math.Pow(2, float64(exp)))) <= n {
					sum += int(math.Pow(2, float64(exp)))
					binary = append(binary, 1)
				} else {
					binary = append(binary, 0)

				}

			}
			break

		} else {
			exp++
			continue

		}

	}
	return binary

}
