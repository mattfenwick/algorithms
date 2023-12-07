package pkg

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var (
	EgTree = &TreeNode{Val: 1,
		Left: &TreeNode{Val: 2,
			Left:  &TreeNode{Val: 4},
			Right: &TreeNode{Val: 5},
		},
		Right: &TreeNode{Val: 3,
			Left: &TreeNode{Val: 6},
		},
	}
)

func BuildTree(length int) *TreeNode {
	if length == 0 { return nil }
	nodes := make([]*TreeNode, length)
	nodes[0] = &TreeNode{Val: 1}
	for i := 1; i < length; i++ {
		val := i + 1
		nodes[i] = &TreeNode{Val: val}
		isLeft := val%2 == 0
		parentIndex := (i - 1) / 2
		if isLeft {
			nodes[parentIndex].Left = nodes[i]
		} else {
			nodes[parentIndex].Right = nodes[i]
		}
	}
	return nodes[0]
}

func spaces(count int) string {
	s := ""
	for i := 0; i < count; i++ {
		s += "   "
	}
	return s
}

func PrintTree(root *TreeNode, depth int) {
	if root == nil {
		return
	}
	PrintTree(root.Right, depth+1)
	fmt.Printf("%s%d\n", spaces(depth), root.Val)
	PrintTree(root.Left, depth+1)
}

func findMaxDepth(root *TreeNode) int {
	depth := 0
	for cur := root; cur != nil; cur = cur.Left {
		fmt.Printf("maxDepth cur: %d\n", cur.Val)
		depth++
	}
	return depth
}

func isPresent(root *TreeNode, n int) bool {
	turns := []bool{}
	for i := n; i > 1; i = i / 2 {
		// fmt.Printf("turn: %d, %d, %t\n", n, i, i % 2 == 0)
		turns = append(turns, i % 2 == 0)
	}
	cur := root
	for j := 0; j < len(turns); j++ {
		// fmt.Printf("%+v, %d, %+v\n", cur, j, turns)
		if turns[len(turns) - j - 1] {
			cur = cur.Left
		} else {
			cur = cur.Right
		}
	}
	return cur != nil
}

func countNodes(root *TreeNode) int {
	if root == nil {
		return 0
	}
	maxDepth := findMaxDepth(root)
	left := 1
	for i := 1; i < maxDepth; i++ {
		left *= 2
	}
	right := left * 2 -1
	for right > left {
		mid := (left + right + 1) / 2
		if isPresent(root, mid) {
			left = mid
		} else {
			if right == mid {
				return left
			}
			right = mid
		}
	}
	return left
}

func CountNodes() {
	fmt.Printf("%d\n", countNodes(EgTree))

	tree := BuildTree(53)
	PrintTree(tree, 0)
	fmt.Printf("nodes: %d\n", countNodes(tree))
	for i := 40; i < 60; i++ {
		fmt.Printf("is present? %d, %t\n", i, isPresent(tree, i))
	}

	for i := 1; i < 650; i++ {
		t := BuildTree(i)
		// PrintTree(t, 0)
		fmt.Printf("nodes: %d, %d\n", i, countNodes(t))
	}
}
