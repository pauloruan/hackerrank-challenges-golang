// Given an array of integers, find the sum of its elements.

// For example, if the array , arr = [1, 2, 3] , 1 + 2 + 3 = 6 so return 6.

package main

import "fmt"

func simpleArraySum(array []int) int {
	var sum int
	for index := 0; index < len(array); index += 1 {
		sum += array[index]
	}
	return sum
}

func main() {
	result := simpleArraySum([]int{10, 20, 30, 40, 50})
	fmt.Println(result)
}
