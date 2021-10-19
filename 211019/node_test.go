package _11019

import "testing"

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


