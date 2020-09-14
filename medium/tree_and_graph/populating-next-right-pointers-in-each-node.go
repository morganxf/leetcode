package tree_and_graph

type Node struct {
	Val   int
	Left  *Node
	Right *Node
	Next  *Node
}

func connect(root *Node) *Node {
	//return connect_1(root)
	return connect_2(root)
}

func connect_1(root *Node) *Node {
	if root == nil {
		return root
	}
	var queue []*Node
	queue = append(queue, root)
	var queueSize int = len(queue)
	for queueSize > 0 {
		for i := 0; i < queueSize; i++ {
			curNode := queue[i]
			if i < queueSize-1 {
				curNode.Next = queue[i+1]
			}
			if curNode.Left != nil {
				queue = append(queue, curNode.Left)
			}
			if curNode.Right != nil {
				queue = append(queue, curNode.Right)
			}
		}
		curQueueSize := len(queue)
		copy(queue, queue[queueSize:])
		queueSize = curQueueSize - queueSize
		queue = queue[:queueSize]
	}
	return root
}

// 利用上次计算结果 动态规划思想
func connect_2(root *Node) *Node {
	if root == nil {
		return nil
	}
	var leftMostNode *Node = root
	for leftMostNode != nil {
		curNode := leftMostNode
		for curNode != nil {
			if curNode.Left != nil {
				curNode.Left.Next = curNode.Right
			}
			if curNode.Next != nil && curNode.Right != nil {
				curNode.Right.Next = curNode.Next.Left
			}
			curNode = curNode.Next
		}
		leftMostNode = leftMostNode.Left
	}
	return root
}
