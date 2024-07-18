package pkg

import (
	"fmt"
	"slices"
)

func targetIndices(nums []int, target int) []int {
	slices.Sort(nums)
	start := findFirst(target, nums)
	indices := []int{}
	if start != -1 {
		for x := start; x < len(nums) && nums[x] == target; x++ {
			indices = append(indices, x)
		}
	}
	return indices
}

func FindTargetIndicesMain() {
	fmt.Printf("2: %+v\n", targetIndices([]int{1, 2, 5, 2, 3}, 2))
	fmt.Printf("3: %+v\n", targetIndices([]int{1, 2, 5, 2, 3}, 3))
	fmt.Printf("5: %+v\n", targetIndices([]int{1, 2, 5, 2, 3}, 5))
}
