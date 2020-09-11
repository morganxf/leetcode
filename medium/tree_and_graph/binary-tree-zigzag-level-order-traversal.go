package tree_and_graph

import (
	"errors"
)

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	var result [][]int
	var q queue
	q.Push(root)
	var line int
	for !q.IsEmpty() {
		var nodes []*TreeNode
		var values []int
		for !q.IsEmpty() {
			node, _ := q.Pop()
			values = append(values, node.Val)
			if node.Left != nil {
				nodes = append(nodes, node.Left)
			}
			if node.Right != nil {
				nodes = append(nodes, node.Right)
			}
		}
		if line%2 == 1 {
			for l, r := 0, len(values)-1; l < r; l, r = l+1, r-1 {
				values[l], values[r] = values[r], values[l]
			}
		}
		line++
		result = append(result, values)
		for _, node := range nodes {
			q.Push(node)
		}
	}
	return result
}

type queue struct {
	nodes []*TreeNode
}

func (q *queue) Push(node *TreeNode) {
	q.nodes = append(q.nodes, node)
}

func (q *queue) Pop() (*TreeNode, error) {
	if len(q.nodes) == 0 {
		return nil, errors.New("empty")
	}
	node := q.nodes[0]
	q.nodes = q.nodes[1:]
	return node, nil
}

func (q *queue) IsEmpty() bool {
	return len(q.nodes) == 0
}
