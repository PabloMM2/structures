package structures

import (
	"fmt"
	"os"
)

type ArrayMethods interface {
	Insert()
	Get()
	Delete()
	Size()
	Print()
}

type Array struct {
	elements []*int
}

func NewArray() ArrayMethods {
	array := &Array{}
	return array
}

func ArrayFunc() {
	array := NewArray()

	for {
		var option int
		fmt.Println()
		fmt.Println("------------------------")
		fmt.Println("| Select and a option: |")
		fmt.Println("| 1: Print             |")
		fmt.Println("| 2: Insert            |")
		fmt.Println("| 3: Get               |")
		fmt.Println("| 4: Delete            |")
		fmt.Println("| 5: Size              |")
		fmt.Println("| 6: Exit              |")
		fmt.Println("------------------------")
		fmt.Scanln(&option)
		switch option {
		case 1:
			array.Print()
		case 2:
			array.Insert()
		case 3:
			array.Get()
		case 4:
			array.Delete()
		case 5:
			array.Size()
		case 6:
			fmt.Println("Good bye")
			os.Exit(0)
		default:
			fmt.Println("Try again!")
		}
	}
}

func (A *Array) Insert() {
	var value int
	fmt.Println("Set the key value >")
	_, err := fmt.Scanln(&value)
	if err != nil {
		return
	}

	A.elements = append(A.elements, &value)
}

func (A *Array) Get() {
	var value int
	fmt.Println("Set the position to be get")
	_, err := fmt.Scanln(&value)
	if err != nil {
		return
	}

	if value < 0 || value >= len(A.elements) {
		fmt.Println("Out of index")
		return
	}

	fmt.Printf("[ %d ]", *A.elements[value])
}

func (A *Array) Delete() {

}

func (A *Array) Size() {
	fmt.Printf("The array size is: %d", len(A.elements))
}

func (A *Array) Print() {
	for _, value := range A.elements {
		fmt.Printf("[%d] ", *value)
	}
}
