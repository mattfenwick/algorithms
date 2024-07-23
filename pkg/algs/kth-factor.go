package algs

import "fmt"

func Reverse[A any](xs []A) []A {
	length := len(xs)
	out := make([]A, length)
	for i, x := range xs {
		out[length-i-1] = x
	}
	return out
}

func GetFactors(n int) []int {
	lows := []int{}
	highs := []int{}
	for i := 1; ; i++ {
		quotient := n / i
		if quotient < i {
			break
		}
		if n%i == 0 {
			lows = append(lows, i)
			if i != quotient {
				highs = append(highs, quotient)
			}
		}
	}
	return append(lows, Reverse(highs)...)
}

func KthFactor(n int, k int) int {
	factors := GetFactors(n)
	if k > len(factors) {
		return -1
	}
	return factors[k-1]
}

func KthFactorMain() {
	pairs := [][2]int{
		{2, 1},
		{2, 2},
		{3, 1},
		{3, 2},
		{3, 3},
		{8, 1},
		{8, 2},
		{8, 3},
		{8, 4},
		{8, 5},
		{144, 6},
		{144, 27},
		{512, 8},
		{521, 1},
		{521, 2},
		{521, 3},
	}
	for _, pair := range pairs {
		n, k := pair[0], pair[1]
		fmt.Printf("%d, %d: %+v, %d\n", n, k, GetFactors(n), KthFactor(n, k))
	}
}
