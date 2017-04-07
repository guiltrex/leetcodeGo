package Leetcode

//O(n) solution
func rob(nums []int) int {
	l := len(nums)
	if l==0 {return 0}
	if l==1 || l==2 || l==3 {return intMax(nums...)}

	m := []int{0, nums[1]}
	for i:= 2; i<len(nums); i++ {
		m[0], m[1] = intMax(m[1], m[0]), m[0]+nums[i]
	}
	res := intMax(m[0], m[1])

	m = []int{nums[0], 0}
	for i:= 2; i<len(nums); i++ {
		m[0], m[1] = intMax(m[1], m[0]), m[0]+nums[i]
	}
	return intMax(res, m[0])
}

//O(nlogn) solution using divide & conquer
func rob(nums []int) int {
	l := len(nums)
	if l==0 {return 0}
	if l==1 || l == 2 || l == 3 {return intMax(nums...)}
	m := make([][]int, 2)
	m[0] = make([]int, 4)
	m[1] = make([]int, 4)
	m[0][0], m[0][1], m[0][2], m[0][3] = subRob(nums[0:l/2])
	m[1][0], m[1][1], m[1][2], m[1][3] = subRob(nums[l/2:l])

	return intMax(  m[0][0]+m[1][0], m[0][0]+m[1][1], m[0][0]+m[1][2], m[0][0]+m[1][3],
		m[0][1]+m[1][0], m[0][1]+m[1][1],
		m[0][2]+m[1][0], m[0][2]+m[1][2],
		m[0][3]+m[1][0])

}

func subRob(nums []int) (int, int, int, int){
	l := len(nums)
	if l==2 {return 0, nums[0], nums[1], 0}
	if l==3 {return nums[1], nums[0], nums[2], nums[0]+nums[2]}

	m := make([][]int, 3)
	m[0] = make([]int, 4)
	m[1] = make([]int, 4)
	m[2] = make([]int, 4)

	m[0][0], m[0][1], m[0][2], m[0][3] = subRob(nums[0:l/2])
	m[1][0], m[1][1], m[1][2], m[1][3] = subRob(nums[l/2:l])

	m[2][0] = intMax(m[0][0]+m[1][0], m[0][0]+m[1][1], m[0][2]+m[1][0])
	m[2][1] = intMax(m[0][1]+m[1][0], m[0][1]+m[1][1], m[0][3]+m[1][0])
	m[2][2] = intMax(m[0][0]+m[1][2], m[0][0]+m[1][3], m[0][2]+m[1][2])
	m[2][3] = intMax(m[0][1]+m[1][2], m[0][1]+m[1][3], m[0][3]+m[1][2])
	return m[2][0],m[2][1],m[2][2],m[2][3]
}

func intMax(slice ...int) int {
	max := slice[0]
	for _, v := range slice {
		if v > max {
			max = v
		}
	}
	return max
}
