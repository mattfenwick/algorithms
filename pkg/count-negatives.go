package pkg

import "fmt"

func countNegatives(grid [][]int) int {
    pos := 0
    x := len(grid[0]) - 1
    for i, row := range grid {
		fmt.Printf("%d %d\n", i, x)
        for x >= 0 && row[x] < 0 {
			fmt.Printf("  %d %d\n", i, x)
            x--
        }
        pos = pos + x + 1
    }
	fmt.Printf("pos: %d\n", pos)
    return len(grid) * len(grid[0]) - pos
}

func CountNegativesMain() {
	grid := [][]int{{4,3,2,-1},{3,2,1,-1},{1,1,-1,-2}, {-1,-1,-2,-3}}
	fmt.Printf("%d\n", countNegatives(grid))
}