// In this challenge, you are required to calculate and print the sum of the elements in an array, keeping in mind that some of those integers may be quite large.

// Complete the aVeryBigSum function in the editor below. It must return the sum of all array elements.

// aVeryBigSum has the following parameter(s): int ar[n]: an array of integers .

// long: the sum of all array elements

package main

import "fmt"

func aVeryBigSum(ar []int64) int64 {
	sum := int64(0)
	for _, number := range ar {
		sum += number
	}
	return sum
}

func main() {
	result := aVeryBigSum([]int64{1000000001, 1000000002, 1000000003, 1000000004, 1000000005})
	fmt.Println(result)
}
