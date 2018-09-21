package main

import "fmt"

type treeNode struct {
	value       int
	left, right *treeNode
}

func (node treeNode) print() {
	fmt.Print(node.value, " ")
}

func (node *treeNode) setValue(value int) {
	if node == nil {
		fmt.Println("Setting value to nil node. Ignored.")
		return
	}
	node.value = value
}

func (node *treeNode) traverse() {
	if node == nil {
		return
	}
	node.left.traverse()
	node.print()
	node.right.traverse()
}
func createNode(value int) *treeNode {
	return &treeNode{value: value}
}

func main() {
	var root treeNode
	//{0 <nil> <nil>}
	fmt.Println(root)

	root = treeNode{value: 3}
	root.left = &treeNode{}
	root.right = &treeNode{5, nil, nil}
	root.right.left = new(treeNode)
	root.left.right = createNode(2)

	root.right.left.setValue(4)
	//0 2 3 4 5
	root.traverse()
	fmt.Println()

	//4
	root.right.left.print()
	fmt.Println()

	root.print()
	//3
	fmt.Println()
	root.setValue(100)

	pRoot := &root
	//100
	pRoot.print()
	fmt.Println()
	pRoot.setValue(200)
	//200
	pRoot.print()
	fmt.Println()

	var ppRoot *treeNode
	//Setting value to nil node. Ignored.
	ppRoot.setValue(200)
	ppRoot = &root
	ppRoot.setValue(300)
	//300
	ppRoot.print()
	fmt.Println()

	nodes := []treeNode{
		{value: 3},
		{},
		{6, nil, &root},
	}
	//[{3 <nil> <nil>} {0 <nil> <nil>} {6 <nil> 0xc0420603e0}]
	fmt.Println(nodes)
}
