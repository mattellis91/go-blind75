package main

/*
Given an integer array nums, find the subarray with the largest sum, and return its sum.

Eg 1:
Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
Output: 6
Explanation: The subarray [4,-1,2,1] has the largest sum 6.

Eg 2:
Input: nums = [1]
Output: 1
Explanation: The subarray [1] has the largest sum 1.

Eg 3:
Input: nums = [5,4,-1,7,8]
Output: 23
Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

Constraints:
1 <= nums.length <= 10^5
-10^4 <= nums[i] <= 10^4
*/


import (
	"fmt"
)

func main() {
	nums := []int{-2,1,-3,4,-1,2,1,-5,4}
	fmt.Println(maxSubArray(nums))
}

func maxSubArray(nums []int) int {
	maxSum := nums[0]
	currentSum := nums[0]
	
	for i := 1; i < len(nums); i++ {
		currentSum = max(nums[i], currentSum + nums[i])
		maxSum = max(maxSum, currentSum)
	}
	
	return maxSum
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}