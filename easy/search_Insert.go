package easy

/*
Given a sorted array and a target value, return the index if the target is found. If not, return the index where it would be if it were inserted in order.

You may assume no duplicates in the array.

Example 1:
Input: [1,3,5,6], 5
Output: 2

Example 2:
Input: [1,3,5,6], 2
Output: 1

Example 3:
Input: [1,3,5,6], 7
Output: 4

Example 4:
Input: [1,3,5,6], 0
Output: 0
*/

//SearchInsert ...
func SearchInsert(nums []int, target int) int {
	// if len(nums) == 0 {
	// 	return 0
	// }
	if nums[len(nums)/2] == target {
		return len(nums) / 2
	} else if nums[len(nums)/2] > target {
		return SearchInsert(nums[0:len(nums)/2], target)
	} else {
		return SearchInsert(nums[len(nums)/2:len(nums)-1], target)
	}
}
