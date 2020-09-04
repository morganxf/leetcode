package linked_list

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	return GetIntersectionNode_1(headA, headB)
}

func GetIntersectionNode_1(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	nodeA := headA
	nodeB := headB
	var isNull bool
	for nodeA != nodeB {
		if nodeA.Next == nil {
			if isNull {
				return nil
			} else {
				nodeA = headB
				isNull = true
			}
		} else {
			nodeA = nodeA.Next
		}
		if nodeB.Next == nil {
			nodeB = headA
		} else {
			nodeB = nodeB.Next
		}
	}
	return nodeA
}