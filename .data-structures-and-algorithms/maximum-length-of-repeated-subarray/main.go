package main

import (
	"fmt"
)

func main() {
	fmt.Println(findLength([]int{1, 2, 3, 2, 1}, []int{3, 2, 1, 4, 7}))
	fmt.Println(findLength([]int{0, 0, 0, 0, 0}, []int{0, 0, 0, 0, 0}))
}

func findLength(nums1 []int, nums2 []int) int {
	max := 0
	for i := 0; i < len(nums1); i++ {
		count := 0
		nextI := i
		nextJ := 0
		for j := 0; j < len(nums2); j++ {
			isCount := false
			if nums1[i] == nums2[j] {
				isCount = true
				nextJ = j
			} else {
				continue
			}

			for {
				if nextI > len(nums1)-1 || nextJ > len(nums2) || nums1[nextI] != nums2[nextJ] {
					break
				}
				count++
				nextI++
				nextJ++
			}

			if isCount == true {
				if count > max {
					max = count
				}
				break
			}

		}
	}
	return max
}
