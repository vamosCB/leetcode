package easy

/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.

思路: 可以考虑传入的字符串切片为一个多维数组
*/

//LongestCommonPrefix 最长公共前缀
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	} else if len(strs) == 1 {
		return strs[0]
	}

	for x := 0; x < len(strs[0]); x++ {
		for y := 1; y < len(strs); y++ {
			if x < len(strs[y]) {
				if strs[0][x] != strs[y][x] {
					return strs[0][0:x]
				}
			} else {
				return strs[0][0:x]
			}
		}
	}
	return strs[0]
}
