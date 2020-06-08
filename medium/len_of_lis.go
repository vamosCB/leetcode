package medium

/*

给定一个无序的整数数组，找到其中最长上升子序列的长度。

示例:

输入: [10,9,2,5,3,7,101,18]
输出: 4
解释: 最长的上升子序列是 [2,3,7,101]，它的长度是 4。
*/

//LengthOfLIS ...
func LengthOfLIS(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	dp[0] = 1
	result := 1
	for i := 0; i < len(nums); i++ {
		maxval := 0
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				maxval = max(maxval, dp[j])
			}
		}
		dp[i] = maxval + 1
		result = max(result, dp[i])
	}
	return result
}

func max(src, dst int) int {
	if src > dst {
		return src
	}
	return dst
}
