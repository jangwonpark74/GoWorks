package addTwoNumbers

import "fmt"

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {

	carry := 0

	head1 := l1
	head2 := l2
	h1 := 0
	h2 := 0
	first := 0
	var root *ListNode
	var prev *ListNode
	for head1 != nil || head2 != nil || carry != 0 {

		node := new(ListNode)
		if head1 == nil {
			h1 = 0
		} else {
			h1 = head1.Val
		}

		if head2 == nil {
			h2 = 0
		} else {
			h2 = head2.Val
		}

		node.Val = (h1 + h2 + carry) % 10

		node.Next = nil
		if first == 0 {
			root = node
			first++
		} else {
			prev.Next = node
		}

		if (h1 + h2 + carry) >= 10 {
			carry = 1
		} else {
			carry = 0
		}

		if head1 != nil {
			head1 = head1.Next
		}
		if head2 != nil {
			head2 = head2.Next
		}
		prev = node
	}

	return root
}
