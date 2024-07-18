package pkg

import (
	"fmt"

	"golang.org/x/exp/slices"
)

type Trie struct {
	Word     *string
	Children [26]*Trie
}

func (t *Trie) Print(depth int) bool {
	if t == nil {
		// fmt.Printf("%s(null)\n", spaces(depth))
		return false
	}
	shouldPrint := false
	if t.Word != nil {
		fmt.Printf("%s%s\n", spaces(depth), *t.Word)
		shouldPrint = true
	}
	for c, child := range t.Children {
		if child.Print(depth + 1) {
			fmt.Printf("%s%da\n", spaces(depth), c)
			shouldPrint = true
		}
	}
	return shouldPrint
}

func Add(t *Trie, word string, index int) *Trie {
	if t == nil {
		t = &Trie{Children: [26]*Trie{}}
	}
	if index == len(word) {
		t.Word = &word
		return t
	}
	key := word[index] - 'a'
	t.Children[key] = Add(t.Children[key], word, index+1)
	return t
}

func GetSubTrie(t *Trie, word string, index int) *Trie {
	if t == nil {
		return nil
	}
	if index == len(word) {
		return t
	}
	key := word[index] - 'a'
	return GetSubTrie(t.Children[key], word, index+1)
}

func FindN(t *Trie, n int) []string {
	if n == 0 || t == nil {
		return nil
	}
	out := []string{}
	if t.Word != nil {
		out = append(out, *t.Word)
		n--
	}
	for _, child := range t.Children {
		if n == 0 {
			break
		}
		next := FindN(child, n)
		n -= len(next)
		out = append(out, next...)
	}
	return out
}

func suggestedProductsTrie(products []string, searchWord string) [][]string {
	t := &Trie{Children: [26]*Trie{}}
	for _, prod := range products {
		t = Add(t, prod, 0)
	}
	// t.Print(0)
	results := [][]string{}
	for i := 1; i <= len(searchWord); i++ {
		subTrie := GetSubTrie(t, searchWord[:i], 0)
		// subTrie.Print(8)
		results = append(results, FindN(subTrie, 3))
		// fmt.Printf("aha %s: %+v\n", searchWord[:i], results[len(results) - 1])
	}
	return results
}

func findFirstPrefix(products []string, prefix string) int {
	comp := func(s string) int {
		sPre := take(s, len(prefix))
		fmt.Printf("is good? %s , %s , %t\n", s, prefix, sPre < prefix)
		if sPre == prefix {
			return 0
		} else if sPre > prefix {
			return 1
		} else {
			return -1
		}
	}
	low, high := 0, len(products)-1
	for {
		mid := (low + high) / 2
		fmt.Printf("  %d, %d, %d\n", low, mid, high)
		compVal := comp(products[mid])
		if low == high {
			if compVal == 0 {
				return mid
			}
			return -1
		}
		if compVal >= 0 {
			high = mid
		} else {
			low = mid + 1
		}
	}
}

func take(s string, n int) string {
	if n > len(s) {
		return s
	}
	return s[:n]
}

func suggestedProducts(products []string, searchWord string) [][]string {
	slices.Sort(products)
	fmt.Printf("products: %+v\n", products)
	out := [][]string{}
	for i := 1; i <= len(searchWord); i++ {
		first := findFirstPrefix(products, searchWord[:i])
		fmt.Printf("looking for %s, found at %d\n", searchWord[:i], first)
		next := []string{}
		if first >= 0 {
			for j := 0; j < 3 && (first+j) < len(products); j++ {
				word := products[first+j]
				fmt.Printf("comparing: %d, %s, %s\n", i, word, searchWord)
				if take(word, i) == searchWord[:i] {
					next = append(next, word)
				}
			}
		}
		out = append(out, next)
	}
	return out
}

func SuggestedProductsMain() {
	if "abc" < "ab" {
		print("yes!\n")
	} else {
		print("no\n")
	}
	fmt.Printf("answer trie: %+v\n", suggestedProductsTrie([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
	fmt.Printf("answer: %+v\n", suggestedProducts([]string{"mobile", "mouse", "moneypot", "monitor", "mousepad"}, "mouse"))
	fmt.Printf("answer: %+v\n", suggestedProducts([]string{"bags", "baggage", "banner", "box", "cloths"}, "bags"))
	fmt.Printf("answer: %+v\n", suggestedProducts([]string{"havana"}, "tatiana"))
}
