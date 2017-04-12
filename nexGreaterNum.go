package Leetcode

//Bruteforce solution.
//Will rewrite O(n) solution after getting user-defined stack in golang (or in other languages)

func nextGreaterElement(findNums []int, nums []int) []int {
	numPosMap := make(map[int]int)
	l := len(nums)
	res := make([]int, len(findNums))

	for i, e := range nums{
		numPosMap[e] = i+1
	}

	nextGreat := func(num int)int{
		for i,_ := numPosMap[num];i<l;i++{
			if num < nums[i]{
				return nums[i]
			}
		}
		return -1
	}

	for i,e := range findNums {
		res[i] = nextGreat(e)
	}
	return res
}

