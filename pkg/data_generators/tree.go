package data_generators

import (
	"hunkevych-philip/algorithms-data-structures/pkg/types/rand"
	"hunkevych-philip/algorithms-data-structures/pkg/types/tree"
	"time"
)

var (
	ExpectedBFSPreOrder = []int{100, 50, 123, 32, 122, 124}
	ExpectedDFSPreOrder = []int{100, 50, 32, 123, 122, 124}
	ExpectedDFSInOrder  = []int{32, 50, 100, 122, 123, 124}
)

func InitTestingTree() *tree.Node {
	myTree := tree.InitNewTree(100)
	myTree.Add(123)
	myTree.Add(50)
	myTree.Add(124)
	myTree.Add(122)
	myTree.Add(32)

	// ^^^^^^^^^^^^^^^
	//			100
	//	       /   \
	//	     50	   123
	// 		/	   /  \
	//	  32	122   124

	// BFS 100 50 123 32 122 124
	// DFS 100 50 32 123 122 124

	return myTree
}

func GetPopulatedTreeOfGivenLength(n int) *tree.Node {
	time.Sleep(time.Second)

	var (
		r      = rand.R()
		myTree = tree.InitNewTree(r.Intn(100))
	)

	for i := 0; i < n; i++ {
		myTree.Add(r.Intn(100))
	}

	return myTree
}
