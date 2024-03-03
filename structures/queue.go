package structures

import (
	"fmt"
)

type QueueMethods interface {
	Print()
	Enqueue()
	Dequeue()
	IsEmpty() bool
	Top()
}

type Queue struct {
	head *NodeQueue
	tail *NodeQueue
}

type NodeQueue struct {
	value *int
	next  *NodeQueue
}

func NewQueue() QueueMethods {
	return &Queue{}
}

func QueueFunc() {

	queue := NewQueue()

	for {
		var option int
		fmt.Println()
		fmt.Println("------------------------")
		fmt.Println("| Select and a option: |")
		fmt.Println("| 1: Print             |")
		fmt.Println("| 2: Enqueue           |")
		fmt.Println("| 3: Dequeue           |")
		fmt.Println("| 4: Is empty ?        |")
		fmt.Println("| 5: Top               |")
		fmt.Println("| 6: Exit              |")
		fmt.Println("------------------------")
		fmt.Scanln(&option)
		switch option {
		case 1:
			queue.Print()
		case 2:
			queue.Enqueue()
		case 3:
			queue.Dequeue()
		case 4:
			queue.IsEmpty()
		case 5:
			queue.Top()
		case 6:
			fmt.Println("Good bye Queue")
			return
		default:
			fmt.Println("Try again!")
		}
	}

}

func (q *Queue) Print() {
	empty := q.IsEmpty()
	if empty {
		fmt.Println("Is empty...")
		return
	}

	node := q.head
	for {
		fmt.Printf("[ %d ] ", *node.value)

		if node.next == nil {
			break
		} else {
			node = node.next
		}
	}

}

func (q *Queue) Enqueue() {
	fmt.Println("Set the node value >")

	var value int
	_, err := fmt.Scan(&value)
	if err != nil {
		fmt.Println("Error value")
		return
	}

	node := &NodeQueue{
		value: &value,
	}

	empty := q.IsEmpty()
	if empty {
		q.head = node
	} else {
		q.tail.next = node
	}

	q.tail = node
}

func (q *Queue) Dequeue() {
	empty := q.IsEmpty()
	if empty {
		fmt.Println("Is empty...")
		return
	}

	head := q.head
	nextNode := head.next
	q.head = nextNode

	if head.next == nil {
		q.tail = nil
	}
}

func (q *Queue) IsEmpty() bool {

	if q.tail == nil && q.head == nil {
		return true
	}
	return false
}

func (q *Queue) Top() {
	empty := q.IsEmpty()
	if empty {
		fmt.Println("Is empty...")
		return
	}

	fmt.Printf("[ %d ]", *q.head.value)
}
