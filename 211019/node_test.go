package _11019

import (
	"fmt"
	"testing"
)

func TestRemoveMiddleNode(t *testing.T) {
	node1 := Node{value: 1}
	node2 := Node{value: 2,next: &node1}
	node3 := Node{value: 3,next: &node2}
	t.Log(node3)
	t.Log(node3.next)
	node := RemoveMiddleNode(&node3)
	t.Log(node)
	t.Log(node.next)
}

func TestReverse(t *testing.T) {
	node1 := Node{value: 1}
	node2 := Node{value: 2,next: &node1}
	node3 := Node{value: 3,next: &node2}
	node := Reverse(&node3)
	t.Log(node)
	t.Log(node.next)
	t.Log(node.next.next)
}

func TestReverseNoRecur(t *testing.T){
	node1 := Node{value: 1}
	node2 := Node{value: 2,next: &node1}
	node3 := Node{value: 3,next: &node2}
	node4 := Node{value: 4,next: &node3}
	node := ReverseNoRecur(&node4)
	t.Log(node)
	t.Log(node.next)
	t.Log(node.next.next)
	t.Log(node.next.next.next)
}

func TestReverseNoRecur2(t *testing.T){
	node1 := Node{value: 1}
	node2 := Node{value: 2,next: &node1}
	node3 := Node{value: 3,next: &node2}
	node4 := Node{value: 4,next: &node3}
	node := ReverseNoRecur2(&node4)
	t.Log(node)
	t.Log(node.next)
	t.Log(node.next.next)
	t.Log(node.next.next.next)
}


func TestIsPalindrome(t *testing.T){
	node1 := Node{value: 1}
	node2 := Node{value: 2,next: &node1}
	node3 := Node{value: 3,next: &node2}
	node4 := Node{value: 4,next: &node3}
	result := IsPalindrome(&node4)
	if result == true{
		t.Fatal("应该为false")
	}
	fmt.Println(node4.String())


	node1 = Node{value: 1}
	node2 = Node{value: 2,next: &node1}
	node3 = Node{value: 2,next: &node2}
	node4 = Node{value: 1,next: &node3}
	result = IsPalindrome(&node4)
	if result ==false{
		t.Fatal("应该为true")
	}
	fmt.Println(node4.String())

	node1 = Node{value: 1}
	node2 = Node{value: 2,next: &node1}
	node3 = Node{value: 3,next: &node2}
	result = IsPalindrome(&node3)
	if result ==true{
		t.Fatal("应该为false")
	}
	fmt.Println(node3.String())

	node1 = Node{value: 1}
	node2 = Node{value: 2,next: &node1}
	node3 = Node{value: 1,next: &node2}
	result = IsPalindrome(&node3)
	if result ==false{
		t.Fatal("应该为true")
	}
	fmt.Println(node3.String())

}


func TestNewLinkList(t *testing.T) {
	node:=NewLinkList([]int{2,5,4,5})
	fmt.Print(node.String())
}

func TestPivot(t *testing.T) {
	node := NewLinkList([]int{9,0,4,5,1})
	newNode := Pivot(node,3)
	fmt.Print(newNode.String())

	node= NewLinkList([]int{0,1,9,3,5})
	newNode = Pivot(node,3)
	fmt.Print(newNode.String())
}


func TestRemoveDuplicate(t *testing.T) {
	node := NewLinkList([]int{1,2,3,3,4,4,2,1,1})
	node2 := RemoveDuplicate(node)
	fmt.Print(node2.String())
}
