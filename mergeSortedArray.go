package Leetcode

func merge(nums1 []int, m int, nums2 []int, n int)  {
	if n == 0 {
		return
	}

	im, in, imn := m-1, n-1, m+n-1

	for ; im >= 0 && in>=0; imn-- {
		if nums2[in] >= nums1[im] {
			nums1[imn] = nums2[in]
			in--
		} else {
			nums1[imn] = nums1[im]
			im--
		}
	}
	if im < 0 {
		for i:= in; i>=0; i-- {
			nums1[i] = nums2[i]
		}
	}
	return
}
