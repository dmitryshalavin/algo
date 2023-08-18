package main

import "fmt"

type ListNode struct {
	Value int
	Next  *ListNode
}

func Reverse(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	r := Reverse(head.Next)
	head.Next.Next = head
	head.Next = nil
	return r
}

func Print(list *ListNode) {
	head := list
	for head != nil {
		fmt.Printf("(%d) -> ", head.Value)
		head = head.Next
	}

	fmt.Println()
}

func main() {
	list := &ListNode{1, &ListNode{2, &ListNode{3, &ListNode{4, &ListNode{5, nil}}}}}

	Print(list)

	list = Reverse(list)

	Print(list)

}
