package main

import (
	"fmt"
)

type Tree struct {
	Value    int
	Children []*Tree
}

var queue []*Tree

func bfs(queue []*Tree) {
	if len(queue) == 0 {
		return
	}

	fmt.Println(queue[0].Value)

	queue = append(queue, queue[0].Children...)

	bfs(queue[1:])

}

func main() {
	tree := &Tree{1, []*Tree{
		{2, []*Tree{
			{3, nil}, {4, []*Tree{{5, nil}}}}},
		{7, []*Tree{{9, nil}}}}}

	queue = []*Tree{tree}
	bfs(queue)
}
