package binarysearch

import (
	"fmt"
)

// receive a number, and slice, make a binary search inside the slice
// return position on the slice, and error if happend
func calMiddle(n1, n2 int) int {
	return (n2 + n1) / 2
}
func binarySearch(number int, arr []int) (int, error) {
	min := arr[0]
	max := arr[len(arr)-1]
	if number < min || number > max {
		return 0, fmt.Errorf(fmt.Sprintf("%v is out of range", number))
	}
	middle := 0
	min = 0
	max = len(arr) - 1
	for min <= max {
		middle = calMiddle(min, max)

		if number < arr[middle] {
			max = middle - 1
		} else if number > arr[middle] {
			min = middle + 1
		} else {
			return middle + 1, nil
		}
	}
	return 0, fmt.Errorf(fmt.Sprintf("%v not found inside the slice", number))
}

func BinarySearch(number int) {

	arr := []int{-20, -12, 0, 1, 3, 5, 6, 11, 22, 33, 55, 88, 100}
	position, err := binarySearch(number, arr)
	fmt.Println("=====================================")
	defer fmt.Println("=====================================")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("Number :%v Position: %v\n", number, position)
	}

}
