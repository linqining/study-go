package _11020

import (
	"fmt"
	"github.com/xcltapestry/xclpkg/algorithm"
)

type Node struct{
	value int
	left *Node
	right *Node
}



func NewTree()*Node{
	node4 := nodeWithVal(4)
	node5 := nodeWithVal(5)
	node2 := nodeWithVal(2)
	node2.left=node4
	node2.right = node5
	node6 := nodeWithVal(6)
	node7 := nodeWithVal(7)
	node3 := nodeWithVal(3)
	node3.left = node6
	node3.right = node7
	node1 := nodeWithVal(1)
	node1.left=node2
	node1.right = node3
	return node1
}

func nodeWithVal(val int)*Node{
	return &Node{value: val}
}

func PreOrderUnRecur(head *Node){
	fmt.Println("pre-order: ")
	if head!=nil{
		stack := algorithm.NewStack()
		stack.Push(head)
		for !stack.Empty(){
			elm := stack.Top()
			err := stack.Pop()
			if err!=nil{
				fmt.Println(err)
				return
			}
			shead := elm.(*Node)
			fmt.Println(shead.value)
			if shead.right!=nil{
				stack.Push(shead.right);
			}
			if shead.left!=nil{
				stack.Push(shead.left)
			}
		}

	}
	fmt.Println("")
}
