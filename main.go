package main

import "fmt"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func main() {
	//fmt.Printf("%v", createNum("9848"))
		
	fmt.Printf("%v", ambiguousCoordinates("-00011-"))
}
