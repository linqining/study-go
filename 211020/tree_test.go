package _11020

import "testing"

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