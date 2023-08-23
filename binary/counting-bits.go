package main

/*
Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n), 
ans[i] is the number of 1's in the binary representation of i.

Example 1:
	Input: n = 2
	Output: [0,1,1]
	Explanation:
	0 --> 0
	1 --> 1
	2 --> 10

Example 2:
	Input: n = 5
	Output: [0,1,1,2,1,2]
	Explanation:
	0 --> 0
	1 --> 1
	2 --> 10
	3 --> 11
	4 --> 100
	5 --> 101
*/

import "fmt"

func main() {
	answer := countBits(5)
	fmt.Printf("%v", answer)
}

func countBits(n int) []int {
	answer := make([]int, n+1)
	for i := 0; i <= n; i++ {
		answer[i] = countOne(i)
	}
	return answer
}

func countOne(n int) int {
	count := 0
	for n > 0 {
		n = n & (n - 1)
		count++
	}
	return count
}