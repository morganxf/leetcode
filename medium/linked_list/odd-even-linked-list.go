package linked_list

func oddEvenList(head *ListNode) *ListNode {
	return OddEvenList(head)
}

func OddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	odd, oddCur := head, head
	even, evenCur := head.Next, head.Next
	for oddCur != nil && evenCur != nil && oddCur.Next !=nil && evenCur.Next != nil{
		oddCur.Next = oddCur.Next.Next
		evenCur.Next = evenCur.Next.Next
		oddCur = oddCur.Next
		evenCur = evenCur.Next
	}
	if oddCur != nil && oddCur.Next != nil {
		oddCur.Next = nil
	}
	if evenCur != nil && evenCur.Next != nil {
		evenCur.Next = nil
	}
	oddCur.Next = even
	return odd
}