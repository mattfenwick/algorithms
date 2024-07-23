package algs

import (
	"fmt"
	"slices"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findTheDistanceValue(arr1 []int, arr2 []int, d int) int {
	out := 0
	slices.Sort(arr2)
	for _, i := range arr1 {
		fmt.Printf("starting: %d\n", i)
		var index int
		if i < arr2[0] {
			index = 0
		} else if i > arr2[len(arr2)-1] {
			index = len(arr2) - 1
		} else {
			index = findfirstByCondition(func(j int) bool {
				fmt.Printf("calc: %d, %d, %d\n", i, j, j-i)
				return j >= i
			}, arr2)
		}
		fmt.Printf("for %d: %d\n", i, index)
		found := false
		if index-1 >= 0 {
			fmt.Printf("type a: %d, %d, %d\n", i, arr2[index-1], d)
			if abs(i-arr2[index-1]) <= d {
				found = true
			}
		}
		if (index >= 0) && (index < len(arr2)) {
			fmt.Printf("type b: %d, %d, %d\n", i, arr2[index], d)
			if abs(i-arr2[index]) <= d {
				found = true
			}
		}
		if !found {
			fmt.Printf("not found: %d\n", i)
			out++
		}
	}
	return out
}

type fdvCase struct {
	arr1 []int
	arr2 []int
	d    int
}

func FindDistanceValueMain() {
	cases := []*fdvCase{
		{[]int{4, 5, 8}, []int{10, 9, 1, 8}, 2},
		{[]int{1, 4, 2, 3}, []int{-4, -3, 6, 10, 20, 30}, 3},
		{[]int{2, 1, 100, 3}, []int{-5, -2, 10, -3, 7}, 6},
		{
			[]int{-803, 715, -224, 909, 121, -296, 872, 807, 715, 407, 94, -8, 572, 90, -520, -867, 485, -918, -827, -728, -653, -659, 865, 102, -564, -452, 554, -320, 229, 36, 722, -478, -247, -307, -304, -767, -404, -519, 776, 933, 236, 596, 954, 464},
			[]int{817, 1, -723, 187, 128, 577, -787, -344, -920, -168, -851, -222, 773, 614, -699, 696, -744, -302, -766, 259, 203, 601, 896, -226, -844, 168, 126, -542, 159, -833, 950, -454, -253, 824, -395, 155, 94, 894, -766, -63, 836, -433, -780, 611, -907, 695, -395, -975, 256, 373, -971, -813, -154, -765, 691, 812, 617, -919, -616, -510, 608, 201, -138, -669, -764, -77, -658, 394, -506, -675, 523, 730, -790, -109, 865, 975, -226, 651, 987, 111, 862, 675, -398, 126, -482, 457, -24, -356, -795, -575, 335, -350, -919, -945, -979, 611},
			37,
		},
		{[]int{-1000, 1, 1000, 2000, 3000}, []int{3, 1000, 1001, 1002}, 5},
		{[]int{-1000, 1, 1000, 2000, 3000}, []int{-1002, -1001, -100, 3}, 5},
		{[]int{-3, 2, -5, 7, 1}, []int{4}, 84},
	}
	for _, c := range cases {
		fmt.Printf("answer: %d:\n  %+v\n  %+v\n  %d\n\n", findTheDistanceValue(c.arr1, c.arr2, c.d), c.arr1, c.arr2, c.d)
	}
}
