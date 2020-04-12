package easy

/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

Example 1:

Input: "()"
Output: true
Example 2:

Input: "()[]{}"
Output: true
Example 3:

Input: "(]"
Output: false
Example 4:

Input: "([)]"
Output: false
Example 5:

Input: "{[]}"
Output: true
*/

//IsValidBrackets 是否为有效的括号
func IsValidBrackets(s string) bool {
	if s == "" {
		return true
	}
	if len(s)%2 == 1 {
		return false
	}
	bracketMap := map[byte]byte{
		'(': ')',
		'{': '}',
		'[': ']',
	}
	var stack []byte
	for index := 0; index < len(s); index++ {
		if _, ok := bracketMap[s[index]]; ok {
			stack = append(stack, s[index])
		} else {
			if len(stack) == 0 {
				return false
			}
			value := stack[len(stack)-1]
			stack = stack[0 : len(stack)-1]
			if bracketMap[value] != s[index] {
				return false
			}
		}
	}
	if len(stack) > 0 {
		return false
	}
	return true
}
