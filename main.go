package main

import (
	"fmt"
	"os"
	"structures/structures"
)

func main() {
	fmt.Println("Welcome!")
	for {
		var option int
		fmt.Println()
		fmt.Println("------------------------")
		fmt.Println("| Select and a option: |")
		fmt.Println("| 1: Array             |")
		fmt.Println("| 2: Stack             |")
		fmt.Println("| 3: Queue             |")
		fmt.Println("| 4: Tree              |")
		fmt.Println("| 6: Exit              |")
		fmt.Println("------------------------")
		fmt.Scanln(&option)
		switch option {
		case 1:
			structures.ArrayFunc()
		case 2:
			structures.StackFunc()
		case 3:
			structures.QueueFunc()
		case 4:
			structures.TreeFunc()
		case 6:
			fmt.Println("Good bye")
			os.Exit(0)
		default:
			fmt.Println("Try again!")
		}
	}
}
