package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {

	// cuối phần dữ liệu thật nums1
	i := m - 1

	// cuối nums2
	j := n - 1

	// cuối mảng kết quả
	k := m + n - 1

	// còn dữ liệu ở cả 2 mảng
	for i >= 0 && j >= 0 {

		// nums1 lớn hơn
		if nums1[i] > nums2[j] {

			nums1[k] = nums1[i]
			i--

		} else {

			nums1[k] = nums2[j]
			j-- 
		}

		k--
	}

	// nếu nums2 còn phần tử
	for j >= 0 {
		nums1[k] = nums2[j]

		j--
		k--
	}
}
 
func main() {
	nums1 := []int{1, 2, 3, 0, 0, 0}
	m := 3
	nums2 := []int{2, 5, 6}
	n := 3
	merge(nums1, m, nums2, n)
	fmt.Println(nums1) // [1 2 2 3 5 6]
}


