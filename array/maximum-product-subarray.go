package main

/*
Given an integer array nums, find a  subarray that has the largest product, and return the product.
The test cases are generated so that the answer will fit in a 32-bit integer.

Eg 1:
Input: nums = [2,3,-2,4]
Output: 6
Explanation: [2,3] has the largest product 6.

Eg 2:
Input: nums = [-2,0,-1]
Output: 0
Explanation: The result cannot be 2, because [-2,-1] is not a subarray.

Constraints:
1 <= nums.length <= 2 * 10^4
-10 <= nums[i] <= 10
*/

import "fmt"

func main() {
	nums := []int{2,3,-2,4}
	fmt.Println(maxProduct(nums))
}

func maxProduct(nums []int) int {
	maxProduct := nums[0]
	currentMaxProduct := nums[0]
	currentMinProduct := nums[0]

	for i := 1; i < len(nums); i++ {
		temp := currentMaxProduct
		currentMaxProduct = max(nums[i], max(nums[i]*currentMaxProduct, nums[i]*currentMinProduct))
		currentMinProduct = min(nums[i], min(nums[i]*temp, nums[i]*currentMinProduct))
		maxProduct = max(maxProduct, currentMaxProduct)
	}

	return maxProduct
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

