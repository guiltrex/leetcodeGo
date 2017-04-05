package Leetcode

func hammingDistance(x int, y int) int {
	hamInt := x ^ y
	hamDist := 0
	for ; hamInt > 0; {
		hamDist += (hamInt & 1)
		hamInt >>= 1
	}
	return hamDist
}
