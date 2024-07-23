package algs

import "fmt"

func Main3Sum() {
	fmt.Printf("answer: %+v\n", threeSum([]int{-1, 1, 0}))
}

func threeSum(nums []int) [][]int {
	negMap, posMap := map[int]int{}, map[int]int{}
	var neg, pos []int
	for _, n := range nums {
		if n < 0 {
			neg = append(neg, n)
			negMap[n]++
		} else {
			pos = append(pos, n)
			posMap[n]++
		}
	}

	var solutions [][]int
	fmt.Printf("%+v\n%+v\n", negMap, posMap)

	for i, ic := range negMap {
		for j := range negMap {
			if j <= i {
				continue
			}
			if _, ok := posMap[-1*(i+j)]; ok {
				solutions = append(solutions, []int{i, j, -1 * (i + j)})
			}
		}
		if ic > 1 {
			if _, ok := posMap[-2*i]; ok {
				solutions = append(solutions, []int{i, i, -2 * i})
			}
		}
	}

	for i, ic := range posMap {
		fmt.Printf("pos: %d\n", i)
		for j := range posMap {
			if j >= i {
				continue
			}
			fmt.Printf("pos inner: %d, %d\n", i, j)
			if _, ok := negMap[-1*(i+j)]; ok {
				solutions = append(solutions, []int{i, j, -1 * (i + j)})
			}
		}
		if ic > 1 {
			if _, ok := negMap[-2*i]; ok {
				solutions = append(solutions, []int{i, i, -2 * i})
			}
		}
	}

	if ct, ok := posMap[0]; ok && ct >= 3 {
		solutions = append(solutions, []int{0, 0, 0})
	}

	return solutions
}
