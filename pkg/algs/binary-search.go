package algs

import (
	"fmt"

	"github.com/mattfenwick/collections/pkg/slice"
)

func binarySearch(xs []int, n int) bool {
	if len(xs) == 0 {
		return false
	}
	low, high := 0, len(xs)-1
	for {
		mid := (low + high) / 2
		// fmt.Printf("%d, %d, %d, %+v\n", low, mid, high, xs)
		if xs[mid] == n {
			return true
		} else if low >= high {
			return false
		} else if xs[mid] > n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
}

func binarySearchExamples() {
	xs := []int{-8, -5, -1, 1, 4, 5, 8, 10, 11, 17, 20, 22, 23, 29}
	ns := slice.Range(-10, 30, 1)
	for _, n := range ns {
		fmt.Printf("%d: %t\n", n, binarySearch(xs, n))
	}
}

// goal: find the lowest index whose element is n.
//
//	  this assumes the array is sorted in increasing order
//
//			if not found, return -1
//		    2, [ 0, 1, 1, 2] =>  3
//		    2, [ 0, 1, 2, 2] =>  2
//			2, [ 1, 2, 2, 3] =>  1
//			2, [ 2, 3, 3, 3] =>  0
//			2, [ 1, 3, 4, 5] => -1
//			2, [-4,-3,-2,-1] => -1
//			2, [ 3, 4, 5, 6] => -1
func findFirst(n int, xs []int) int {
	low, high := 0, len(xs)-1
	found := false
	index := len(xs)
	for {
		mid := (low + high) / 2
		// fmt.Printf("%d, %d, %d\n", low, mid, high)
		if xs[mid] == n && mid < index {
			found = true
			index = mid
		}
		if high <= low {
			break
		} else if xs[mid] >= n {
			high = mid - 1
		} else {
			low = mid + 1
		}
	}
	if !found {
		return -1
	}
	return index
}

// goal: find last index whose element is n
//
//	if not found, return -1
func findLast(n int, xs []int) int {
	low, high := 0, len(xs)-1
	index := -1
	for {
		mid := (low + high) / 2
		// fmt.Printf("find last: %d, %d, %d\n", low, mid, high)
		if xs[mid] == n && mid > index {
			index = mid
		}
		if low >= high {
			break
		} else if xs[mid] <= n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return index
}

// check whether n is present
func binarySearchIsPresent(n int, xs []int) bool {
	low, high := 0, len(xs)-1
	for {
		mid := (low + high) / 2
		if xs[mid] == n {
			return true
		}
		if high <= low {
			break
		} else if xs[mid] < n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	return false
}

// goal: find first index whose element is greater than n
//
//	if not found, return -1
func findFirstGreater(n int, xs []int) int {
	low, high := 0, len(xs)-1
	found, index := false, len(xs)
	for {
		mid := (low + high) / 2
		if xs[mid] > n && mid < index {
			found = true
			index = mid
		}
		if high <= low {
			break
		} else if xs[mid] <= n {
			low = mid + 1
		} else {
			high = mid - 1
		}
	}
	if !found {
		return -1
	}
	return index
}

func findFirstClean(n int, xs []int) int {
	low, high := 0, len(xs)-1
	for {
		mid := (low + high) / 2
		if high <= low {
			if xs[mid] == n {
				return mid
			}
			return -1
		} else if xs[mid] > n {
			high = mid - 1
		} else if xs[mid] == n {
			high = mid
		} else {
			low = mid + 1
		}
	}
}

func BinarySearchMain() {
	if false {
		egs := [][]int{
			{0, 1, 1, 2},
			{0, 1, 2, 2},
			{1, 2, 2, 3},
			{2, 3, 3, 3},
			{1, 3, 4, 5},
			{-4, -3, -2, -1},
			{3, 4, 5, 6},
			{-8, -7, -4, 0, 0, 0, 3, 3, 4, 7, 9},
			{-8, -7, -4, 0, 0, 2, 3, 3, 4, 7, 9},
			{-8, -7, -4, 0, 0, 0, 2, 3, 4, 7, 9},
			{-8, -7, -4, 0, 0, 3, 3, 4, 7, 9},
			{-8, -7, -4, 0, 0, 2, 3, 3, 4, 7, 9},
			{-8, -7, -4, 0, 0, 2, 2, 2, 2, 3, 3, 4, 7, 9},
			{-8, -7, -4, 0, 0, 2, 2, 3, 3, 4, 7, 9},
		}
		for _, e := range egs {
			fmt.Printf("%d, %+v: %t, %d, %d, %d, %d\n", 2, e, binarySearchIsPresent(2, e), findFirst(2, e), findLast(2, e), findFirstGreater(2, e), findFirstClean(2, e))
		}
	} else {
		binarySearchExamples()
	}
}
