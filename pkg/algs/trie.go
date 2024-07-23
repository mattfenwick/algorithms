package algs

// type Trie struct {
// 	Char     string
// 	Children map[string]*Trie
// }

// func (t *Trie) Add(xs string) {
// 	if len(xs) == 0 {
// 		return
// 	}
// 	fst, rest := xs[0], xs[1:]
// 	child, ok := t.Children[fst]
// 	if !ok {
// 		child = &Trie{Char: fst, Children: map[string]*Trie{}}
// 		t.Children[fst] = child
// 	}
// 	child.Add(rest)
// }

// func (t *Trie) Contains(xs string) bool {
// 	if len(xs) == 0 {
// 		return true
// 	}
// 	fst, rest := xs[0], xs[1:]
// 	if child, ok := t.Children[fst]; ok {
// 		return child.Contains(rest)
// 	}
// 	return false
// }
