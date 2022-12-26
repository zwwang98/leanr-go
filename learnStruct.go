package main

import "fmt"

func main() {
	useStruct()
}

/*
1. Go仅支持封装，不支持继承和多态
2. Go没有class，只有struct
3. 对于Go，不需要知道变量是分配在栈上还是堆上
   对于Java、C++的情况我不熟悉
4. 给struct定义的方法，有一个receiver，相当于Java里的this，然后通过实例加"."直接调用这个方法
   Go里所有参数都是传值，所以如果想通过函数修改receiver里的参数，必须将receiver明确为pointer
5. nil指针也可以作为方法的receiver来调用方法，所以需要在函数实现里加一个nil check if statement，比如在下面的setValue里：
	func (node *treeNode) setValue(v int) {
		if node == nil {
			fmt.Println("Setting value to nil node. Ignored.")
			return
		}
		node.value = v
	}
6. 选择 值receiver && pointer receiver
  a. 如果要改变receiver内的值，使用pointer
  b. 如果receiver过大，为了避免在值传递时需要对较大的receiver进行拷贝而影响性能，推荐使用pointer
  c. 一致性 - 如果有指针接受者，最好都用指针接受者
7. 比较 值receiver && pointer receiver
  a. 值receiver是go特有的
  b. 不论一个函数的receiver是值还是pointer，调用时都是一样的，都是该类型变量加"."来调用方法
*/

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Println(node.value)
}

func (node *treeNode) setValue(v int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = v
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}

// 自定义工厂函数
// C++里，返回局部变量会报错，go不会
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func useStruct() {
	var root treeNode

	root = treeNode{value: 3}
	root.left = &treeNode{}
	// treeNode{}返回的不是指针，是实例，需要手动取地址
	root.right = &treeNode{5, nil, nil}
	// 不管是指针还是实例，都直接用"."访问
	// new返回一个指针
	root.right.left = new(treeNode)

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	fmt.Println(nodes)

	root.right.right = createNode(19)
	fmt.Println(root)

	fmt.Println("\nBefore change root's value, root:")
	root.print()
	fmt.Println("Change root's value to 100")
	root.setValue(100)
	fmt.Println("After change root's value, root:")
	root.print()

	fmt.Println("\nIn-order traversal")
	root.traverse()
}
