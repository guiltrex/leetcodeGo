package Leetcode

//Two loop solution
//Will add ONE loop solution
func longestPalindrome(s string) int {
	cm := make(map[rune]int)
	for _, v := range s{
		cm[v] += 1
	}
	res, odd := 0, 0
	for _, v := range cm{
		res += v - v%2
		odd = odd | v%2
	}
	return res + odd
}
