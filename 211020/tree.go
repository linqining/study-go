package _11020

import (
	"fmt"
	"github.com/xcltapestry/xclpkg/algorithm"
	"math"
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

func PosOrderUnRecur2(head *Node){
	fmt.Println("pos-order2: ")
	if head!=nil{
		stack := algorithm.NewStack()
		stack.Push(head)
		var c *Node
		for !stack.Empty(){
			elm := stack.Top()
			c = elm.(*Node)
			if c.left!=nil && head != c.left && head!=c.right{
				stack.Push(c.left)
			}else if c.right!=nil && head!=c.right{
				stack.Push(c.right)
			}else{
				elm := stack.Top()
				stack.Pop()
				fmt.Println(elm)
				head = c
			}
		}
	}
	fmt.Println("")
}


func BiggestSubBST(head *Node) *Node{
	record :=make([]int,3)
	return postOrder(head, record)
}

// 这里注意传入的是引用类型
func postOrder(head *Node,record []int)*Node{
	if head==nil{
		record[0]=0
		record[1] = math.MaxInt32
		record[2] = math.MinInt32
		return nil
	}
	value := head.value
	fmt.Println(value)
	left := head.left
	right := head.right

	lBST := postOrder(left,record)
	lSize := record[0]
	lMin := record[1]
	lMax := record[2]
	rBST := postOrder(right,record)
	rSize := record[0]
	rMin := record[1]
	rMax := record[2]
	if lMin>rMin{
		record[1] = rMin
	}else{
		record[1] = lMin
	}
	if rMax>lMax{
		record[2] = rMax
	}else{
		record[2] = lMax
	}
	if left == lBST && right == rBST && lMax<value && value < rMin{
		record[0] = lSize + rSize+1
		return head
	}
	if lSize>rSize{
		record[0] =lSize
	}else{
		record[0] = rSize
	}
	if lSize>rSize{
		return lBST
	}else{
		return rBST
	}
}