package Leetcode

import "strings"

//not efficient enough, require refactoring
func reverseString(s string) string {
	res := make([]byte, len(s))
	for i := 0; i<len(s); i++ {
		res[len(s)-1-i] = s[i]
	}
	return string(res)
}
