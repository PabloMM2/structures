package structures

import (
	"fmt"
)

type StackMethods interface {
	Push()
	Pop()
	IsEmpty()
	Top()
	Print()
}

type Stack struct {
	head *Node
}

type Node struct {
	value *int
	next  *Node
}

func NewStack() StackMethods {
	stack := &Stack{}
	return stack
}

func StackFunc() {

	stack := NewStack()

	for {
		var option int
		fmt.Println()
		fmt.Println("------------------------")
		fmt.Println("| Select and a option: |")
		fmt.Println("| 1: Print             |")
		fmt.Println("| 2: Push              |")
		fmt.Println("| 3: Pop               |")
		fmt.Println("| 4: Is empty ?        |")
		fmt.Println("| 5: Top               |")
		fmt.Println("| 6: Exit              |")
		fmt.Println("------------------------")
		fmt.Scanln(&option)
		switch option {
		case 1:
			stack.Print()
		case 2:
			stack.Push()
		case 3:
			stack.Pop()
		case 4:
			stack.IsEmpty()
		case 5:
			stack.Top()
		case 6:
			fmt.Println("Good bye")
			return
		default:
			fmt.Println("Try again!")
		}
	}
}

func (S *Stack) Push() {
	var value int
	fmt.Println("Set the node value >")
	_, err := fmt.Scanln(&value)
	if err != nil {
		return
	}

	node := &Node{
		value: &value,
		next:  S.head,
	}

	S.head = node
}

func (S *Stack) Pop() {

	if S.head == nil {
		fmt.Println("The Stack is empy")
		fmt.Println("\"\"")
		return
	}

	S.head = S.head.next
}

func (S *Stack) IsEmpty() {
	if S.head == nil {
		fmt.Println("Is empty")
	} else {
		fmt.Println("Is no't empty")
	}
}

func (S *Stack) Top() {

	if S.head == nil {
		fmt.Println("The Stack is empy")
		fmt.Println("\"\"")
		return
	}

	fmt.Printf("[ %d ]", *S.head.value)
}

func (S *Stack) Print() {
	node := S.head

	if node == nil {
		fmt.Println("The Stack is empy")
		fmt.Println("\"\"")
		return
	}

	for {
		fmt.Printf("[ %d ]", *node.value)
		fmt.Println()

		if node.next == nil {
			break
		} else {
			node = node.next
		}
	}
}
