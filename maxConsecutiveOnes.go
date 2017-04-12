package Leetcode

func findMaxConsecutiveOnes(nums []int) int {
	oneNum := 0
	maxNum := 0
	for i:=0; i<len(nums); i++{
		if nums[i]==0 {
			maxNum = max(maxNum, oneNum)
			oneNum = 0
		} else {
			oneNum++
		}
	}
	return max(maxNum, oneNum)
}

func max(i, j int) int{
	if i>j {
		return i
	}
	return j
}
