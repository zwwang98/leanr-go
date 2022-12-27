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

/*
通过embedding来拓展已有类型
 1. embedding 相当于把 tree.Node 里的所有成员变量都拉出来作为 embedTreeNode 的直接成员变量和方法，
    所以我们可以直接通过 Left, Right 来访问原本属于 tree.Node 的成员变量
 2. 这样的好处是，可以省掉组合写法中为了访问到更内层的成员变量和方法而加入的 .node
*/
type embedTreeNode struct {
	*tree.Node
}

func (myNode *embedTreeNode) postOrderTraversal() {
	if myNode == nil || myNode.Node == nil {
		return
	}

	left := embedTreeNode{myNode.Left}
	right := embedTreeNode{myNode.Right}

	left.postOrderTraversal()
	right.postOrderTraversal()
	myNode.Print()
}

/*
使用了通过 composition 来拓展 node 的 myTreeNode
*/
func mainUsingComposition() {
	fmt.Println("\n使用了通过 composition 来拓展 node 的 myTreeNode")

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

/*
使用了通过 embedding 来拓展 node 的 embedTreeNode
*/
func mainUsingEmbedding() {
	fmt.Println("\n使用了通过 embedding 来拓展 node 的 embedTreeNode")

	root := embedTreeNode{&tree.Node{Value: 3}}
	root.Left = &tree.Node{}
	// tree.TreeNode{}返回的不是指针，是实例，需要手动取地址
	root.Right = &tree.Node{5, nil, nil}
	// 不管是指针还是实例，都直接用"."访问
	// new返回一个指针
	root.Right.Left = new(tree.Node)

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
	root.postOrderTraversal()
}

func main() {
	mainUsingComposition()
	mainUsingEmbedding()
}
