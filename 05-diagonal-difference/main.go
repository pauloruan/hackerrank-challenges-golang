// Given a square matrix, calculate the absolute difference between the sums of its diagonals.
// For example, the square matrix  is shown below:
// 1 2 3
// 4 5 6
// 9 8 9  
// The left-to-right diagonal = . The right to left diagonal = . Their absolute difference is .

// Function description
// Complete the  function in the editor below.
// diagonalDifference takes the following parameter:
// int arr[n][m]: an array of integers
// Return
// int: the absolute diagonal difference

package main

import (
	"fmt"
)

func diagonalDifference(arr [][]int32) int32 {
	var leftDiagonal, rightDiagonal, absoluteValue int32
	for index := 0; index < len(arr); index += 1 {
		leftDiagonal += arr[index][index]
		rightDiagonal += arr[index][len(arr) - index - 1]
	}
	absoluteValue = leftDiagonal - rightDiagonal
	if absoluteValue < 0 { absoluteValue = -absoluteValue}
	return absoluteValue
}

func main() {
	result := diagonalDifference([][]int32{ {1, 2, 3}, {4, 5, 6}, {9, 8, 9} })
	fmt.Println(result)
}
