package structures

import (
	"fmt"
)

type LinkedListMethods interface {
	Print()
	Push()
	Pop()
	Top()
	Last()
	Remove()
	Add()
	IsEmpty() bool
}

type LinkedList struct {
	Head   *LinkedListNode
	Tail   *LinkedListNode
	length uint
}

type LinkedListNode struct {
	Value *int
	Next  *LinkedListNode
}

func NewLinkedList() LinkedListMethods {
	return &LinkedList{}
}

func LinkedListFunc() {

	linkedL := NewLinkedList()

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
		fmt.Println("| 6: Last              |")
		fmt.Println("| 7: Remove            |")
		fmt.Println("| 8: Add               |")
		fmt.Println("| 9: Exit              |")
		fmt.Println("------------------------")
		fmt.Scanln(&option)
		switch option {
		case 1:
			linkedL.Print()
		case 2:
			linkedL.Push()
		case 3:
			linkedL.Pop()
		case 4:
			linkedL.IsEmpty()
		case 5:
			linkedL.Top()
		case 6:
			linkedL.Last()
		case 7:
			linkedL.Remove()
		case 8:
			linkedL.Add()
		case 9:
			fmt.Println("Good bye")
			return
		default:
			fmt.Println("Try again!")
		}
	}
}

func (ll *LinkedList) Print() {
	if ll.IsEmpty() {
		return
	}

	node := ll.Head

	for {
		if node == ll.Head {
			fmt.Printf("Head [ %d ] -> ", *node.Value)
		} else if node == ll.Tail {
			fmt.Printf("Tail [ %d ] -> ", *node.Value)
		} else {
			fmt.Printf("[ %d ] -> ", *node.Value)
		}

		if node.Next == nil {
			break
		} else {
			node = node.Next
		}
	}
}

func (ll *LinkedList) Push() {
	fmt.Println("Set the node value >")

	var value int
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println("Error value")
		return
	}

	node := &LinkedListNode{
		Value: &value,
	}

	if ll.IsEmpty() {
		ll.Head = node
	} else {
		ll.Tail.Next = node
	}

	ll.Tail = node
	ll.length += 1
}

func (ll *LinkedList) Pop() {
	if ll.IsEmpty() {
		return
	}

	node := ll.Head

	if ll.Head == ll.Tail {
		node = nil
		ll.Head = nil
		ll.Tail = nil
		ll.length = 0
		return
	}

	for {
		if node.Next == ll.Tail {
			node.Next = nil
			ll.Tail = node
			ll.length--
			break
		} else {
			node = node.Next
		}
	}
}

func (ll *LinkedList) IsEmpty() bool {

	if ll.Tail == nil && ll.Head == nil {
		fmt.Println("Is Empty")
		return true
	}
	fmt.Println("Is Not Empty")
	return false
}

func (ll *LinkedList) Top() {

	if ll.IsEmpty() {
		return
	}

	fmt.Printf("[ %d ]", *ll.Head.Value)
}

func (ll *LinkedList) Last() {

	if ll.IsEmpty() {
		return
	}

	fmt.Printf("[ %d ]", *ll.Tail.Value)
}

func (ll *LinkedList) Remove() {

	if ll.IsEmpty() {
		return
	}

	fmt.Println("Set the position to remove >")
	var position int
	_, err := fmt.Scan(&position)
	if err != nil {
		fmt.Println("Error value")
		return
	}

	if position < 0 || position > int(ll.length)-1 {
		fmt.Println("Not valid position")
		return
	}

	node := ll.Head

	if position == int(ll.length)-1 {
		ll.Pop()
		return
	} else if position == 0 {
		if ll.Head == ll.Tail {
			ll.Tail = nil
		}
		ll.Head = node.Next
		node = nil
		ll.length--
		return
	}

	counter := 0
	for {
		if position == counter+1 {
			next := node.Next
			node.Next = next.Next
			next = nil
			ll.length--
			return
		} else {
			node = node.Next
			counter++
		}
	}
}

func (ll *LinkedList) Add() {

	fmt.Println("Set the position to Add >")
	var position int
	_, err := fmt.Scan(&position)
	if err != nil {
		fmt.Println("Error value")
		return
	}

	if position < 0 || position >= int(ll.length) {
		fmt.Println("Not valid position")
		return
	}

	if position == int(ll.length)-1 || ll.IsEmpty() {
		ll.Push()
		return
	}

	fmt.Println("Set the node value >")
	var value int
	_, errVal := fmt.Scan(&value)
	if errVal != nil {
		fmt.Println("Error value")
		return
	}

	newNode := &LinkedListNode{
		Value: &value,
	}

	if position == 0 {
		newNode.Next = ll.Head
		ll.Head = newNode
		ll.length++
		return
	}

	node := ll.Head
	counter := 0
	for {
		if position == counter+1 {
			next := node.Next
			newNode.Next = next
			node.Next = newNode
			ll.length++
			return
		} else {
			node = node.Next
			counter++
		}
	}
}
