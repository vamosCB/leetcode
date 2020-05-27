package easy

/*
Implement strStr().

Return the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

Example 1:

Input: haystack = "hello", needle = "ll"
Output: 2
Example 2:

Input: haystack = "aaaaa", needle = "bba"
Output: -1
*/

//StrStr ...
func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	if (haystack == "" && needle == "") || len(needle) > len(haystack) {
		return -1
	}
	slow := 0

	for slow < len(haystack) {
		if haystack[slow] == needle[0] {
			if slow+len(needle) > len(haystack) {
				return -1
			}
			if haystack[slow:slow+len(needle)] == needle {
				return slow
			}
		}
		slow++
	}
	if slow == len(haystack) {
		return -1
	}
	return slow
}
