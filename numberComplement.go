package Leetcode

//Also able to handle negative input
func findComplement(num int) int {
	var mask int = ^0
	for ; mask & num != 0; mask <<= 1 {}
	return num ^ (^mask)
}