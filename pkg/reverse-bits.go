package pkg

import (
	"fmt"
)

func ReverseBitsMain() {
	a := uint32(4322327)
	out := uint32(0)
	for i := 0; i < 32; i++ {
		out += ((a >> i) & 1) << (32 - i - 1)
		fmt.Printf("%d, %d\n    %s, %s\n", a, out, bits(a), bits(out))
	}
	fmt.Printf("%d, %d, %d, %d, %d\n", a, a>>1, a>>2, a>>3, a>>4)
	fmt.Printf("%d, %d, %d, %d, %d\n", a&1, a&2, a&4, a&8, a&16)
	fmt.Printf("%d, %d, %d, %d, %d\n", 1<<4, 1<<3, 1<<2, 1<<1, 1<<0)
}

func bits(n uint32) string {
	out := ""
	for i := 0; i < 32; i++ {
		out += format((n >> (32 - i - 1)) & 1)
	}
	return out
}

func format(i uint32) string {
	if i == 0 {
		return "0"
	} else if i == 1 {
		return "1"
	}
	panic("OOPS")
}
