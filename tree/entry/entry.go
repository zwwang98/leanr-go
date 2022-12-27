package main

import (
	"fmt"
	"learn-go/tree"
)

type myTreeNode struct {
	node *tree.Node
}

/*
https://www.geeksforgeeks.org/inheritance-in-golang/
通过composition拓展已有类型
*/
func (myNode *myTreeNode) postOrderTraversal() {
	if myNode == nil || myNode.node == nil {
		return
	}

	left := myTreeNode{myNode.node.Left}
	right := myTreeNode{myNode.node.Right}

	left.postOrderTraversal()
	right.postOrderTraversal()
	myNode.node.Print()
}

func main() {
	var root tree.Node

	root = tree.Node{Value: 3}
	root.Left = &tree.Node{}
	// tree.TreeNode{}返回的不是指针，是实例，需要手动取地址
	root.Right = &tree.Node{5, nil, nil}
	// 不管是指针还是实例，都直接用"."访问
	// new返回一个指针
	root.Right.Left = new(tree.Node)

	nodes := []tree.Node{
		{Value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.Right.Right = tree.CreateNode(19)
	fmt.Println(root)

	fmt.Println("\nBefore change root's Value, root:")
	root.Print()
	fmt.Println("Change root's Value to 100")
	root.SetValue(100)
	fmt.Println("After change root's Value, root:")
	root.Print()

	fmt.Println("\nIn-order traversal")
	root.InOrderTraverse()

	fmt.Println("\nPost-order traversal")
	myRoot := myTreeNode{&root}
	myRoot.postOrderTraversal()
}
