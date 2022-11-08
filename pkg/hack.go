package pkg

import "fmt"

func HackMain() {
	arr := []int{1, 2, 3, 4, 5}
	fst, rest := arr[0], arr[1:]
	fmt.Printf("%d, %+v\n", fst, rest)
	for i, c := range "abc" {
		fmt.Printf("%d %+v %T\n", i, c, c)
	}
}
