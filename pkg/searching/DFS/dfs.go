package DFS

import (
	"hunkevych-philip/algorithms-data-structures/pkg/types/stack"
	"hunkevych-philip/algorithms-data-structures/pkg/types/tree"
)

// GetDfsBySlicePreOrder Depth-First Search with the help of slice (PreOrder)
// it is much faster and has less memory complexity than the stack
func GetDfsBySlicePreOrder(n *tree.Node) []int {
	res := make([]int, 0)

	manager := make([]*tree.Node, 0)
	currentNode := n
	for {
		// moving by left row till nil val
		if currentNode != nil {
			manager = append(manager, currentNode)
			// --
			res = append(res, currentNode.Value)
			// --
			currentNode = currentNode.Left
			continue
		}
		// left is nil -> time to find first right node starting from the last item
		for i := len(manager) - 1; i >= 0; i-- {
			if manager[i].Right != nil {
				currentNode = manager[i].Right
				manager = manager[:i] // get rid of processed nodes
				break
			}
			if i == 0 {
				return res
			}
		}
	}
}

type DfsRecursionHelper struct {
	NodeManager []*tree.Node
	Index       *tree.Node
}

// GetDfsBySliceRPreOrder the same as above but with the help of recursion
func GetDfsBySliceRPreOrder(n *tree.Node) []int {
	return helper(n, make([]int, 0))
}

func helper(n *tree.Node, res []int) []int {
	res = append(res, n.Value)

	if n.Left != nil {
		res = helper(n.Left, res)
	}

	if n.Right != nil {
		res = helper(n.Right, res)
	}

	return res
}

//GetDfsBySliceInOrder Depth-First Search with the help of slice (InOrder)
func GetDfsBySliceInOrder(n *tree.Node) []int {
	res := make([]int, 0)

	manager := make([]*tree.
		Node, 0)
	currentNode := n
	for {
		// moving by left row till nil val
		if currentNode != nil {
			manager = append(manager, currentNode)
			currentNode = currentNode.Left
			continue
		}
		// left is nil -> time to find first right node starting from the last item
		// otherwise will append each item to a result set
		for i := len(manager) - 1; i >= 0; i-- {
			res = append(res, manager[i].Value)

			if manager[i].Right != nil {
				currentNode = manager[i].Right
				manager = manager[:i] // get rid of processed nodes
				break
			}
			if i == 0 {
				return res
			}
		}
	}
}

//GetDfsByStackPreOrder Depth-First Search with the help of stack (PreOrder)
func GetDfsByStackPreOrder(n *tree.Node) []int {
	res := make([]int, 0)

	s := stack.Stack{}
	c := n
	for {
		// moving by left row till nil val
		if c != nil {
			s.Push(c)
			// --
			res = append(res, c.Value)
			// --
			c = c.Left
			continue
		}
		// left is nil -> time to find first right node starting from the last item
		for {
			if s.Length == 0 {
				return res
			}

			c = s.Peek().(*tree.
				Node).Right
			s.Pop()
			if c != nil {
				break
			}
		}
	}
}
