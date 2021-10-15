/*
 * @lc app=leetcode.cn id=4 lang=golang
 *
 * [3] 无重复字符的最长子串
 */
//@lc code=start
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	l1 := len(nums1)
	l2 := len(nums2)
	l := l1 + l2
	isOdd := 0
	if l%2 == 1 {
		isOdd = 1
	}
	//corner case
	if l == 0 {
		return 0.00000
	}
	//双指针
	Ptr1 := 0
	Ptr2 := 0
	resSlice := make([]int, 0, l)
	res := float64(0)
	for Ptr1 < l1 && Ptr2 < l2 {
		if nums1[Ptr1] <= nums2[Ptr2] {
			resSlice = append(resSlice, nums1[Ptr1])
			Ptr1++
		} else {
			resSlice = append(resSlice, nums2[Ptr2])
			Ptr2++
		}
	}
	if Ptr1 == l1 && Ptr2 < l2 {
		resSlice = append(resSlice, nums2[Ptr2:]...)
		Ptr2 = l2
	} else if Ptr1 < l1 && Ptr2 == l2 {
		resSlice = append(resSlice, nums1[Ptr1:]...)
		Ptr1 = l1
	}
	if isOdd == 1 {
		index := int(math.Floor(float64(l / 2)))
		res = float64(resSlice[index])
	} else {
		res = float64((resSlice[l/2-1] + resSlice[l/2])) / 2.0
	}
	res = math.Floor(res*100000) / 100000
	fmt.Println(resSlice)
	return res

}

// @lc code=end

