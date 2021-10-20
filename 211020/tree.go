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

func InOrderUnRecur(head *Node){
	fmt.Println("in-order: ")
	if head!=nil{
		stack := algorithm.NewStack()
		for !stack.Empty()||head!=nil{
			if head!=nil{
				stack.Push(head)
				head = head.left
			}else{
				elm := stack.Top()
				err := stack.Pop()
				if err!=nil{
					fmt.Println("打印错误",err)
				}
				head = elm.(*Node)
				fmt.Println(head.value)
				head = head.right
			}

		}
	}
}

func PosOrderUnRecur(head *Node){
	fmt.Println("pos-order: ")
	if head!=nil{
		s1 := algorithm.NewStack()
		s2:= algorithm.NewStack()
		s1.Push(head)
		for !s1.Empty(){
			elm := s1.Top()
			s1.Pop()
			head = elm.(*Node)
			s2.Push(head);
			if head.left!=nil{
				s1.Push(head.left);
			}
			if head.right!=nil{
				s1.Push(head.right)
			}
		}
		for !s2.Empty(){
			elm := s2.Top()
			s2.Pop()
			fmt.Println(elm)
		}
	}
	fmt.Println("")
}