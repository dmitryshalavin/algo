package main

import "fmt"

type Tree struct {
	Value int
	Left  *Tree
	Right *Tree
}

func dfs(root *Tree) {
	if root == nil {
		return
	}
	fmt.Println(root.Value)
	dfs(root.Left)
	dfs(root.Right)
}

func main() {
	tree := &Tree{1, &Tree{2, &Tree{3, nil, nil}, &Tree{4, &Tree{5, nil, nil}, nil}}, &Tree{7, nil, &Tree{9, nil, nil}}}
	dfs(tree)
}
