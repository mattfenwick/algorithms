package pkg

import "fmt"

func daysForCapacity(weights []int, capacity int) int {
	day := 0
	for i := 0; i < len(weights); day++ {
		if weights[i] > capacity {
			return -1
		}
		// fmt.Printf("day %d\n", day)
		for remaining := capacity; i < len(weights) && weights[i] <= remaining; i++ {
			// fmt.Printf("  %d, %d\n", remaining, i)
			remaining -= weights[i]
		}
	}
	return day
}

func shipWithinDays(weights []int, days int) int {
	low, high := 1, 25000000
	for {
		mid := (low + high) / 2
		daysNeeded := daysForCapacity(weights, mid)
		fmt.Printf("%d, %d, %d, %d\n", low, mid, high, daysNeeded)
		if low == high {
			if daysNeeded < 0 {
				panic("AHHHHH")
			}
			return low
		} else if daysNeeded < 0 || daysNeeded > days {
			low = mid + 1
		} else {
			high = mid
		}
	}
}

type ShipWithinCase struct {
	Weights []int
	Days    int
}

func ShipWithinDaysMain() {
	cases := []*ShipWithinCase{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
		{[]int{3, 2, 2, 4, 1, 4}, 3},
		{[]int{1, 2, 3, 1, 1}, 4},
	}
	for _, c := range cases {
		fmt.Printf("%+v, days %d: need capacity %d\n\n", c.Weights, c.Days, shipWithinDays(c.Weights, c.Days))
	}
}
