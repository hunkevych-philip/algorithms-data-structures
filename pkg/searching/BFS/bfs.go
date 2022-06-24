package BFS

import (
	"hunkevych-philip/algorithms-data-structures/pkg/types/queue"
	"hunkevych-philip/algorithms-data-structures/pkg/types/tree"
)

// GetBfsByQueuePreOrder Breadth-First Search with the help of a queue
func GetBfsByQueuePreOrder(n *tree.Node) []int {
	res := make([]int, 0)

	q := queue.Queue{}
	q.Enqueue(n)

	for q.Length > 0 {
		if q.First.Val.(*tree.Node).Left != nil {
			q.Enqueue(q.First.Val.(*tree.Node).Left)
		}
		if q.First.Val.(*tree.Node).Right != nil {
			q.Enqueue(q.First.Val.(*tree.Node).Right)
		}

		res = append(res, q.First.Val.(*tree.Node).Value)
		q.Unshift()
	}

	return res
}

// GetBfsByQueuePreOrderRecursion Breadth-First Search with the help of a queue and recursion
// it's faster than slice method, but memory consumption is higher
func GetBfsByQueuePreOrderRecursion(n *tree.Node) []int {
	q := queue.Queue{}
	q.Enqueue(n)

	return helper(q)
}

func helper(current queue.Queue) (res []int) {
	next := queue.Queue{}

	for current.Length > 0 {
		if current.First.Val.(*tree.Node).Left != nil {
			next.Enqueue(current.First.Val.(*tree.Node).Left)
		}
		if current.First.Val.(*tree.Node).Right != nil {
			next.Enqueue(current.First.Val.(*tree.Node).Right)
		}

		res = append(res, current.First.Val.(*tree.Node).Value)
		current.Unshift()
	}

	if next.Length != 0 {
		res = append(res, helper(next)...)
	}

	return
}

// GetBfsBySlicePreOrder Breadth-First Search with the help of a slice (PreOrder) (level by level solution)
// it's slower than queue method, but memory consumption is lower
func GetBfsBySlicePreOrder(n *tree.Node) []int {
	res := make([]int, 0)

	currentLevel := []*tree.Node{n} // put head as a starting point

	for {
		lowerLevel := make([]*tree.Node, 0)
		for _, currentNode := range currentLevel {
			res = append(res, currentNode.Value)

			if currentNode.Left != nil {
				lowerLevel = append(lowerLevel, currentNode.Left)
			}
			if currentNode.Right != nil {
				lowerLevel = append(lowerLevel, currentNode.Right)
			}
		}
		if len(lowerLevel) == 0 {
			break
		}

		currentLevel = lowerLevel
	}

	return res
}
