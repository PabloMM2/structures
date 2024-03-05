package structures

import (
	"fmt"
)

type TreeMethods interface {
	AddNode()
	InOrder(*NodeTree)
	GetRoot() *NodeTree
}

type Tree struct {
	root *NodeTree
}

type NodeTree struct {
	value *int
	Left  *NodeTree
	Right *NodeTree
}

func NewTree() TreeMethods {
	return &Tree{}
}

func TreeFunc() {
	tree := NewTree()

	for {
		var option int
		fmt.Println()
		fmt.Println("------------------------")
		fmt.Println("| Select and a option: |")
		fmt.Println("| 1: Print             |")
		fmt.Println("| 2: Add node          |")
		fmt.Println("| 6: Exit              |")
		fmt.Println("------------------------")
		fmt.Scanln(&option)
		switch option {
		case 1:
			//queue.Print()
		case 2:
			tree.AddNode()
		case 3:
			root := tree.GetRoot()
			tree.InOrder(root)
		case 4:
			//queue.IsEmpty()
		case 5:
			//queue.Top()
		case 6:
			fmt.Println("Good bye Queue")
			return
		default:
			fmt.Println("Try again!")
		}
	}
}

func (t *Tree) AddNode() {
	var value int
	fmt.Println("Set The node Value")
	_, err := fmt.Scan(&value)

	if err != nil {
		fmt.Println("Wrong value")
		return
	}

	node := &NodeTree{
		value: &value,
	}

	searchNode := t.root

	if searchNode == nil {
		t.root = node
		return
	}

	for {
		if *searchNode.value > *node.value {
			if searchNode.Left == nil {
				searchNode.Left = node
				break
			}
			searchNode = searchNode.Left
			continue
		}

		if searchNode.Right == nil {
			searchNode.Right = node
			break
		}

		searchNode = searchNode.Right
		continue
	}
}

func (t *Tree) InOrder(node *NodeTree) {
	if node == nil {
		return
	}

	t.InOrder(node.Left)
	fmt.Println(*node.value)
	t.InOrder(node.Right)
}

func (t *Tree) GetRoot() *NodeTree {
	return t.root
}
