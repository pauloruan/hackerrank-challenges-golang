// Complete the plusMinus function in the editor below.

// plusMinus has the following parameter(s):

// int arr[n]: an array of integers
// Print the ratios of positive, negative and zero values in the array. Each value should be printed on a separate line with  digits after the decimal. The function should not return a value.

package main

import "fmt"

func plusMinus(arr []int32) {
	var positives, negatives, zeros int32
	for _, value := range arr {
		if value > 0 {
			positives++
		} else if value < 0 {
			negatives++
		} else {
			zeros++
		}
	}
	fmt.Printf("%.6f\n", float32(positives)/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(negatives)/float32(len(arr)))
	fmt.Printf("%.6f\n", float32(zeros)/float32(len(arr)))
}

func main() {
	arr := []int32{-4, 3, -9, 0, 4, 1}
	plusMinus(arr)
}
