package Leetcode

func intersection(nums1 []int, nums2 []int) []int {
	res := []int{}
	if len(nums1)==0 || len(nums2)==0 { return res }

	rMap := make(map[int]int)
	kMap := make(map[int]struct{})
	for _, v:= range nums1 {
		kMap[v] = struct{}{}
	}

	for _, k:= range nums2 {
		if _, ok := kMap[k]; ok {
			rMap[k]+=1
		}
	}

	for k, _:= range rMap {
		res = append(res, k)
	}

	return res
}
