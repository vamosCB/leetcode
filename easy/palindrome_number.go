package easy

/*
Determine whether an integer is a palindrome. An integer is a palindrome when it reads the same backward as forward

Example 1:
	Input: 121
	Output: true

Example 2:
	Input: -121
	Output: false
	Explanation: From left to right, it reads -121. From right to left, it becomes 121-. Therefore it is not a palindrome.

*/

//IsPalindrome 回文数
func IsPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	origin := x
	var temp int
	for x != 0 {
		temp = temp*10 + x%10
		x = x / 10
	}
	return temp == origin
}
