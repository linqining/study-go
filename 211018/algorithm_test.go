package _11018

import "testing"

func TestStack_Reverse(t *testing.T) {
	stack := Stack{}
	for _,item := range [5]int{1,2,3,4,5}{
		stack.Push(item)
	}
	t.Log(stack)
	stack.Reverse()
	t.Log(stack)
}


func TestStack_Sort(t *testing.T) {
	stack := Stack{}
	for _,item := range [5]int{59,2,43,4,5}{
		stack.Push(item)
	}
	t.Log(stack)
	stack.Sort()
	t.Log(stack)
}