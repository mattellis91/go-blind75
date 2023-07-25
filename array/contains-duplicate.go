package main

/*
Given an integer array nums, return true if any value appears at least twice in the array, and return false if every element is distinct.

https://leetcode.com/problems/contains-duplicate/

Eg 1
Input: nums = [1,2,3,1]
Output: true

Eg 2
Input: nums = [1,2,3,4]
Output: false

Eg 3
Input: nums = [1,1,1,3,3,4,3,2,4,2]
Output: true

Constraints:
1 <= nums.length <= 10^5
-109 <= nums[i] <= 10^9
*/

import "fmt"

func main() {
	answer := containsDuplicate([]int{1,2,3,1})
	fmt.Printf("%v", answer)
}

func containsDuplicate(nums []int) bool {
	lookup := make(map[int]bool)

	for _, v := range nums {
		if _, ok := lookup[v]; ok {
			return true
		} else {
			lookup[v] = true
		}
	}

	return false
}