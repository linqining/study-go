package _11020

import (
	"fmt"
	"testing"
)

func TestPreOrderUnRecur(t *testing.T) {
	node := NewTree()
	PreOrderUnRecur(node)
}

func TestInOrderUnRecur(t *testing.T) {
	node := NewTree()
	InOrderUnRecur(node)
}

func TestPosOrderUnRecur(t *testing.T) {
	node := NewTree()
	PosOrderUnRecur(node)
}

func TestPosOrderUnRecur2(t *testing.T) {
	node := NewTree()
	PosOrderUnRecur2(node)
}


func TestBiggestSubBST(t *testing.T) {
	node0 :=nodeWithVal(0)
	node3 := nodeWithVal(3)
	node1 := nodeWithVal(1)
	node1.left=node0
	node1.right = node3

	node2 := nodeWithVal(2)
	node5 := nodeWithVal(5)
	node4 := nodeWithVal(4)
	node4.left = node2
	node4.right = node5

	node11 := nodeWithVal(11)
	node15 := nodeWithVal(15)
	node14 := nodeWithVal(14)
	node14.left = node11
	node14.right = node15

	node10 := nodeWithVal(10)
	node10.left = node4
	node10.right = node14

	node20:= nodeWithVal(20)
	node16 := nodeWithVal(16)
	node13 := nodeWithVal(13)
	node13.left = node20
	node13.right = node16

	node12 := nodeWithVal(12)
	node12.left = node10
	node12.right = node13

	node6 := nodeWithVal(6)
	node6.left = node1
	node6.right = node12

	node := BiggestSubBST(node6)
	fmt.Println(node)
}