package Leetcode

//Also able to handle negative input
func findComplement(num int) int {
	unum := uint32(num)
	var xorInt uint32 = 1
	var shiftNum uint32 = unum >> 1
	for ; shiftNum != 0; shiftNum >>= 1 {
		xorInt = xorInt << 1 + 1
	}
	return int(unum ^ xorInt)
}