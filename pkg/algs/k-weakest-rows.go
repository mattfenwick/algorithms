package algs

import (
	"fmt"
	"slices"
)

func findLastDecreasing(n int, xs []int) int {
	low, high := 0, len(xs)-1
	index := -1
	for {
		mid := (low + high) / 2
		if xs[mid] == n && mid > index {
			index = mid
		}
		if low >= high {
			break
		} else if xs[mid] >= n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return index
}

type row struct {
	strength int
	index    int
}

func kWeakestRows(mat [][]int, k int) []int {
	rows := make([]*row, len(mat))
	for i, matRow := range mat {
		rows[i] = &row{strength: findLastDecreasing(1, matRow) + 1, index: i}
	}
	slices.SortStableFunc(rows, func(a *row, b *row) int {
		if b.strength != a.strength {
			return a.strength - b.strength
		}
		return a.index - b.index
	})
	for _, r := range rows {
		fmt.Printf("%d, %d\n", r.index, r.strength)
	}
	out := make([]int, k)
	for i := 0; i < k; i++ {
		out[i] = rows[i].index
	}
	return out
}

func KWeakestRowsMain() {
	matrix := [][]int{
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 0},
		{1, 0, 0, 0, 0},
		{1, 1, 0, 0, 0},
		{1, 1, 1, 1, 1},
	}
	for _, row := range matrix {
		fmt.Printf("last, dec: %+v, %d\n", row, findLastDecreasing(1, row))
	}
	fmt.Printf("k weakest: %+v\n", kWeakestRows(matrix, 3))
}
