package Leetcode

func maxSubArray(nums []int) int {
	if len(nums)==0 {return 0}
	cur := nums[0]
	res := nums[0]
	for i:=1;i<len(nums);i++{
		cur = intMax(cur + nums[i], nums[i])
		res = intMax(res, cur)
	}
	return res
}

func intMax(i, j int) int {
	if i >= j {
		return i
	} else {
		return j
	}
}
