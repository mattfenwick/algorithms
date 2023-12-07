package pkg

import "slices"

func targetIndices(nums []int, target int) []int {
	slices.Sort(nums)
}

func FindTargetIndicesMain() {
	targetIndices([]int{1,2,5,2,3}, 2)
}