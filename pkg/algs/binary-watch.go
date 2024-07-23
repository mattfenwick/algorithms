package algs

import "fmt"

func MainBinaryWatch() {
	for _, x := range readBinaryWatch(8) {
		fmt.Printf("solution: %s\n", x)
	}
}

var binaryWatchBits = [][2]int{
	{0, 1},
	{0, 2},
	{0, 4},
	{0, 8},
	{0, 16},
	{0, 32},
	{1, 0},
	{2, 0},
	{8, 0},
	{4, 0},
}

func readBinaryWatch(turnedOn int) []string {
	return helper(0, 0, turnedOn, [2]int{0, 0})
}

func add(xs [2]int, ys [2]int) [2]int {
	return [2]int{
		xs[0] + ys[0],
		xs[1] + ys[1],
	}
}

func helper(index int, turnedOn int, turnedOnGoal int, total [2]int) []string {
	fmt.Printf("%d, %d, %d, %+v\n", index, turnedOn, turnedOnGoal, total)
	if total[0] >= 12 || total[1] >= 60 {
		return nil
	}
	if turnedOn == turnedOnGoal {
		return []string{fmt.Sprintf("%d:%02d", total[0], total[1])}
	}
	if index == len(binaryWatchBits) {
		return nil
	}
	out := helper(index+1, turnedOn+1, turnedOnGoal, add(binaryWatchBits[index], total))
	out = append(out, helper(index+1, turnedOn, turnedOnGoal, total)...)
	return out
}
