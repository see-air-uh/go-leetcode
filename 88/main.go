package main

import "fmt"

func merge(nums1 []int, m int, nums2 []int, n int) {
	var mCounter = m - 1
	var nCounter = n - 1
	var idx = m + n - 1
	for idx >= 0 {
		if nCounter < 0 {
			idx = -1
			break
		}
		if mCounter >= 0 && nums1[mCounter] >= nums2[nCounter] {
			nums1[idx] = nums1[mCounter]
			mCounter = mCounter - 1
		} else if nCounter >= 0 {
			nums1[idx] = nums2[nCounter]
			nCounter = nCounter - 1
		}
		idx = idx - 1
	}
	return
}

func main() {
	m := 3
	n := 3
	nums1 := make([]int, m+n)   // Create a slice with a length of m+n
	copy(nums1, []int{1, 2, 3}) // Copy the initial values into nums1
	nums2 := []int{2, 5, 6}

	// [1,2,3,0,0,0]
	merge(nums1, m, nums2, n)
	fmt.Println(nums1)

}
