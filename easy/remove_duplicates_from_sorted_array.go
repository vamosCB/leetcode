package easy

import "fmt"

/*
Given a sorted array nums, remove the duplicates in-place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this by modifying the input array in-place with O(1) extra memory.

Example 1:

Given nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively.

It doesn't matter what you leave beyond the returned length.
Example 2:

Given nums = [0,0,1,1,1,2,2,3,3,4],

Your function should return length = 5, with the first five elements of nums being modified to 0, 1, 2, 3, and 4 respectively.

It doesn't matter what values are set beyond the returned length.
*/

//RemoveDuplicates 从有序数组中移除重复元素
func RemoveDuplicates(nums []int) int {
	if nums == nil || len(nums) == 0 {
		return 0
	}
	cur := 0
	for index := 1; index < len(nums); index++ {
		if nums[cur] != nums[index] {
			cur++
			nums[cur] = nums[index]
		}

	}
	fmt.Printf("nums = %v\n", nums)
	//这里的返回结果其实后面有累赘, 但是有这一句话It doesn't matter what values are set beyond the returned length.
	return cur + 1
}
