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
	if len(nums) == 0 {
		return 0
	}
	tail := len(nums)
	head := 0

	for head <= tail {
		mid := head + (tail-head)/2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			head = mid + 1
		} else {
			tail = mid - 1
		}
	}
	return head
}
