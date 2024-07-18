package pkg

import "fmt"

type lruArray struct {
	maxSize int
	elems   []int
	head    int
	lowest  int
}

func newLruArray(maxSize int) *lruArray {
	elems := make([]int, maxSize)
	for i := 0; i < maxSize; i++ {
		elems[i] = -1
	}
	return &lruArray{
		maxSize: maxSize,
		elems:   elems,
		head:    0,
		lowest:  -1,
	}
}

func (l *lruArray) add(i int) {
	l.head = (l.head + 1) % l.maxSize
	l.lowest = l.elems[l.head]
	l.elems[l.head] = i
}

func getMax(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func almostString(answerKey string, val rune, k int) int {
	a := newLruArray(k)
	max := 0
	for x, t := range answerKey {
		if t != val {
			a.add(x)
		}
		curr := x - a.lowest
		// fmt.Printf("  curr: %s, %s, %d, %d, %d, %d\n", string(val), string(t), a.lowest, a.head, x, curr)
		max = getMax(max, curr)
	}
	return max
}

func maxConsecutiveAnswers(answerKey string, k int) int {
	return max(almostString(answerKey, 'T', k), almostString(answerKey, 'F', k))
}

func MaxConsecutiveAnswersMain() {
	fmt.Printf("%t, %t, %t\n", 'T' == 'T', 'T' == 'F', "F"[0] == 'F')
	fmt.Printf("%d\n\n", almostString("TFFT", 'T', 1))
	fmt.Printf("%d\n\n", almostString("TFFT", 'F', 1))
	fmt.Printf("%d\n\n", almostString("TTFTTFTT", 'T', 1))
	fmt.Printf("%d\n\n", almostString("TTFTTFTT", 'F', 1))
	fmt.Printf("%d\n", almostString("TTFTTFTT", 'T', 2))
	fmt.Printf("%d\n", almostString("TTFTTFTT", 'F', 2))
	fmt.Printf("%d\n", almostString("TTFTTFTTFTTFTT", 'T', 3))
}
