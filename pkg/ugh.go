package pkg

import "fmt"

func UghMain() {
	for n := -1; n < 8; n++ {
		fmt.Printf("%d\n", searchInsert([]int{1, 3, 5, 6}, n))
	}
}

func searchInsert(nums []int, target int) int {
	return searchInsertHelper(0, len(nums), nums, target)
}

func searchInsertHelper(low int, high int, nums []int, target int) int {
	fmt.Printf("helper: %d, %d, %+v, %d\n", low, high, nums, target)
	if low >= len(nums) {
		return low
	}
	if high-low <= 1 {
		if nums[low] >= target {
			return low
		} else {
			return high
		}
	}
	mid := (low + high) / 2
	if nums[mid] == target {
		return mid
	} else if nums[mid] > target {
		return searchInsertHelper(low, mid, nums, target)
	} else {
		return searchInsertHelper(mid+1, high, nums, target)
	}
}
