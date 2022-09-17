package main

import (
	"fmt"
)

// findMaxLength return the maximum length of a subarray that appears in both arrays
func findMaxLength(arr1 []int, arr2 []int) int {
	max := 0
	l1, l2 := len(arr1), len(arr2)
	dp := make([][]int, l1+1)
	for k := range dp {
		dp[k] = make([]int, l2+1)
	}
	for i := l1 - 1; i >= 0; i-- {
		for j := l2 - 1; j >= 0; j-- {
			if arr1[i] == arr2[j] {
				dp[i][j] = dp[i+1][j+1] + 1
				if max < dp[i][j] {
					max = dp[i][j]
				}
			}
		}
	}

	return max
}

func main() {

	fmt.Println(findMaxLength([]int{0, 0, 0, 0}, []int{0, 0, 0, 0}))

	fmt.Println(findMaxLength([]int{0, 2, 0, 0}, []int{0, 4, 2, 0}))

	fmt.Println(findMaxLength([]int{1, 2, 3, 0}, []int{1, 2, 3, 0}))

	fmt.Println(findMaxLength([]int{0, 2, 0, 3}, []int{0, 5, 0, 1}))

	fmt.Println(findMaxLength([]int{1, 2, 3, 4}, []int{5, 6, 7, 8}))
}
