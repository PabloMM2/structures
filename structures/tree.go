package structures

import (
	"fmt"
)

type TreeMethods interface {
	AddNode()
	InOrder(*NodeTree)
	PreOrder(*NodeTree)
	PostOrder(*NodeTree)
	GetRoot() *NodeTree
	Remove(node *NodeTree, value *int) *NodeTree
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
		fmt.Println("| 3: InOrder           |")
		fmt.Println("| 4: PreOrder          |")
		fmt.Println("| 5: PostOrder         |")
		fmt.Println("| 6: Remove            |")
		fmt.Println("| 7: Exit              |")
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
			root := tree.GetRoot()
			tree.PreOrder(root)
		case 5:
			root := tree.GetRoot()
			tree.PostOrder(root)
		case 6:
			root := tree.GetRoot()
			var value int

			fmt.Println("Set The node Value to find or remove")
			_, err := fmt.Scan(&value)
			if err != nil {
				fmt.Println("Wrong value")
			}

			tree.Remove(root, &value)
		case 7:
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
	fmt.Printf("( %d )", *node.value)
	t.InOrder(node.Right)
}

func (t *Tree) PreOrder(node *NodeTree) {
	if node == nil {
		return
	}

	fmt.Printf("( %d )", *node.value)
	t.PreOrder(node.Left)
	t.PreOrder(node.Right)
}

func (t *Tree) PostOrder(node *NodeTree) {
	if node == nil {
		return
	}

	t.PostOrder(node.Left)
	t.PostOrder(node.Right)

	fmt.Printf("( %d )", *node.value)
}

func (t *Tree) GetRoot() *NodeTree {
	return t.root
}

func (t *Tree) Remove(node *NodeTree, value *int) *NodeTree {

	if node == nil {
		return node
	}

	if *value < *node.value {
		node.Left = t.Remove(node.Left, value)
	}

	if *value > *node.value {
		node.Right = t.Remove(node.Right, value)
	}

	if *value == *node.value {

		if node.Left == nil && node.Right == nil {
			fmt.Println("No children")
			node = nil
		} else if node.Left == nil {
			fmt.Println("No left children")
			temp := node
			node = node.Right

			if temp != nil {
				temp = nil
			}

		} else if node.Right == nil {
			fmt.Println("No Right children")
			temp := node
			node = node.Left

			if temp != nil {
				temp = nil
			}
		}

	}

	return node
}

func (t *Tree) SearchNode() *NodeTree {
	var value int

	fmt.Println("Set The node Value to find or remove")
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println("Wrong value")
		return nil
	}

	node := t.root
	for {

		if *node.value == value {
			fmt.Println("Node Found")
			return node
		}

		if value < *node.value {
			if node.Left == nil {
				fmt.Println("Node not found")
				return nil
			}
			node = node.Left
		}

		if value > *node.Right.value {
			if node.Right == nil {
				fmt.Println("Node not found")
				return nil
			}
			node = node.Right
		}
	}
}
