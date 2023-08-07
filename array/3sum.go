package main

/*
Given an integer array nums, return all the triplets [nums[i], nums[j], nums[k]] such that i != j, i != k, and j != k, and nums[i] + nums[j] + nums[k] == 0.
Notice that the solution set must not contain duplicate triplets.

Eg 1:
Input: nums = [-1,0,1,2,-1,-4]
Output: [[-1,-1,2],[-1,0,1]]
Eplanation:
nums[0] + nums[1] + nums[2] = (-1) + 0 + 1 = 0.
nums[1] + nums[2] + nums[4] = 0 + 1 + (-1) = 0.
nums[0] + nums[3] + nums[4] = (-1) + 2 + (-1) = 0.
The distinct triplets are [-1,0,1] and [-1,-1,2].
Notice that the order of the output and the order of the triplets does not matter.

Eg 2:
Input: nums = [0,1,1]
Output: []
Explanation: The only possible triplet does not sum up to 0.

Eg 3
Input: nums = [0,0,0]
Output: [[0,0,0]]
Explanation: The only possible triplet sums up to 0.
*/

import (
	"sort"
	"fmt"
)

func main() {
	nums := []int{-1,0,1,2,-1,-4}
	answer := threeSum(nums)
	fmt.Println(answer)
}

func threeSum(nums []int) [][]int {
    res := [][]int{} 
    
    sort.Ints(nums) 
    
    for i := 0; i < len(nums)-2; i++ { 
        if(i == 0 || (i > 0 && nums[i] != nums[i-1])) { 
            low := i+1 
            high := len(nums)-1
            sum := 0-nums[i]
            
            for (low < high) { 
                if(nums[low] + nums[high] == sum) { 
                    res = append(res, []int{nums[i], nums[low], nums[high]}) 
                    for (low < high && nums[low] == nums[low+1]) { 
						low++ 
					}      
                    for (low < high && nums[high] == nums[high-1]) { 
						high-- 
					}
                    low++
                    high--
                } else if (nums[low] + nums[high] > sum) {
                    high--
                } else {
                    low++
                }
            }
        }    
    }
    
    return res
    
}