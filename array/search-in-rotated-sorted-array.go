package main

/*
There is an integer array nums sorted in ascending order (with distinct values).

Prior to being passed to your function, nums is possibly rotated at an unknown pivot index k (1 <= k < nums.length) such that the resulting array is [nums[k], nums[k+1], ..., nums[n-1], nums[0], nums[1], ..., nums[k-1]] (0-indexed). For example, [0,1,2,4,5,6,7] might be rotated at pivot index 3 and become [4,5,6,7,0,1,2].

Given the array nums after the possible rotation and an integer target, return the index of target if it is in nums, or -1 if it is not in nums.

You must write an algorithm with O(log n) runtime complexity.

Eg 1:
Input: nums = [4,5,6,7,0,1,2], target = 0
Output: 4

Eg 2:
Input: nums = [4,5,6,7,0,1,2], target = 3
Output: -1

Eg 3:
Input: nums = [1], target = 0
Output: -1

Constraints:
1 <= nums.length <= 5000
-10^4 <= nums[i] <= 10^4
All values of nums are unique.
nums is guaranteed to be rotated at some pivot.
-10^4 <= target <= 10^4
*/

import "fmt"

func main() {
	nums := []int{4,5,6,7,0,1,2}
	target := 0
	fmt.Println(search(nums, target))
}

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	
	low := 0
	high := len(nums) - 1
	
	for low <= high {
		mid := (low + high) / 2
		
		if nums[mid] == target {
			return mid
		}
		
		if nums[low] <= nums[mid] {
			if target >= nums[low] && target <= nums[mid] {
				high = mid - 1
			} else {
				low = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[high] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}
	}
	
	return -1
}