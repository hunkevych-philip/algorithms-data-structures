package tree

func InitNewTree(val int) *Node {
	rootNode := &Node{
		Value: val,
	}
	return rootNode
}

// Add adds a value to a tree
func (n *Node) Add(val int) {
	c := n
	for {
		if val < c.Value {
			if c.Left != nil {
				c = c.Left
				continue
			}
			c.Left = &Node{
				Value: val,
			}
			return
		} else {
			if c.Right != nil {
				c = c.Right
				continue
			}
			c.Right = &Node{
				Value: val,
			}
			return
		}
	}
}
