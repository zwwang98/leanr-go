package tree

func (node *Node) InOrderTraverse() {
	if node == nil {
		return
	}
	node.Left.InOrderTraverse()
	node.Print()
	node.Right.InOrderTraverse()
}
