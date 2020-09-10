package tree_and_graph

import "errors"

func inorderTraversal(root *TreeNode) []int {
	//return InorderTraversal_1(root)
	return InorderTraversal_2(root)
}

func InorderTraversal_1(root *TreeNode) []int {
	var datas []int
	inorderTraversalHelper(&datas, root)
	return datas
}

func inorderTraversalHelper(datas *[]int, root *TreeNode) {
	if root == nil {
		return
	}
	inorderTraversalHelper(datas, root.Left)
	*datas = append(*datas, root.Val)
	inorderTraversalHelper(datas, root.Right)
}

type stack struct {
	nodes []*TreeNode
	size  int
}

func (s *stack) Push(node *TreeNode) {
	s.nodes = append(s.nodes, node)
	s.size++
}

func (s *stack) Pop() (*TreeNode, error) {
	if s.size == 0 {
		return nil, errors.New("empty")
	}
	s.size--
	node := s.nodes[s.size]
	s.nodes = s.nodes[:s.size]
	return node, nil
}

func InorderTraversal_2(root *TreeNode) []int {
	var result []int
	var s stack
	cur := root
	var err error
	for {
		if cur != nil {
			s.Push(cur)
			cur = cur.Left
		} else {
			cur, err = s.Pop()
			if err != nil {
				break
			}
			result = append(result, cur.Val)
			cur = cur.Right
		}
	}
	return result
}
