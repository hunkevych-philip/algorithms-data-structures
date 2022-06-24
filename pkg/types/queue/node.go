package queue

type Node struct {
	Val  interface{}
	Next *Node
}
