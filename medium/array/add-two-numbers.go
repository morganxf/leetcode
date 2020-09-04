package array

type ListNode struct {
	Val int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	return AddTwoNumbers_1(l1, l2)
}

func AddTwoNumbers_1(l1 *ListNode, l2 *ListNode) *ListNode {
	var third int
	var result *ListNode = &ListNode{}
	head := result
	for l1 != nil && l2 != nil {
		tmp := &ListNode{}
		tmp.Val = l1.Val + l2.Val + third
		third = tmp.Val / 10
		tmp.Val = tmp.Val % 10

		result.Next = tmp
		l1 = l1.Next
		l2 = l2.Next
		result = result.Next
	}
	var tmpNode *ListNode
	if l1 != nil {
		tmpNode = l1
	} else {
		tmpNode = l2
	}
	for tmpNode != nil {
		tmp := &ListNode{}
		tmp.Val = tmpNode.Val + third
		third = tmp.Val / 10
		tmp.Val = tmp.Val % 10

		result.Next = tmp
		tmpNode = tmpNode.Next
		result = result.Next
	}
	if third != 0 {
		tmp := &ListNode{Val: third}
		result.Next = tmp
		result = result.Next
	}
	return head.Next
}