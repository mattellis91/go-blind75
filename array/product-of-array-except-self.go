package main

/*
Given an integer array nums, return an array answer such that answer[i] is equal to the product of all the elements of nums except nums[i].

The product of any prefix or suffix of nums is guaranteed to fit in a 32-bit integer.

You must write an algorithm that runs in O(n) time and without using the division operation.

Eg 1:
Input: nums = [1,2,3,4]
Output: [24,12,8,6]

Eg 2:
Input: nums = [-1,1,0,-3,3]
Output: [0,0,9,0,0]

Constrains:
2 <= nums.length <= 10^5
-30 <= nums[i] <= 30
*/

import "fmt"

func main() {
	answer := productExceptSelf([]int{1,2,3,4})
	fmt.Printf("%v", answer)
}

func productExceptSelf(nums []int) []int {
	result := make([]int, len(nums))
	
	result[0] = 1
	for i := 1; i < len(nums); i++ {
		result[i] = result[i-1] * nums[i-1]
	}
	
	right := 1
	for i := len(nums)-1; i >= 0; i-- {
		result[i] = result[i] * right
		right *= nums[i]
	}
	
	return result
}