package Leetcode

func rob(nums []int) int {
	if len(nums) < 1 {return 0}
	if len(nums) == 1 {return nums[0]}

	m := nums[0:2]

	for i:=2;i<len(nums);i++{
		m[0], m[1] = intMax(m[1], m[0]), m[0] + nums[i]
	}
	return intMax(m[0], m[1])
}

func intMax(i, j int) int{
	if i>=j {
		return i
	} else {
		return j
	}
}
