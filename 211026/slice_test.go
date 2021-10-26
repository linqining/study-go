package _11026

import "testing"

func TestSlice(t *testing.T){
	// it has sufficient capacity, the destination is resliced to accommodate the
	// new elements. If it does not, a new underlying array will be allocated.
	// Append returns the updated slice. It is therefore necessary to store the
	// result of append, often in the variable holding the slice itself:
	//	slice = append(slice, elem1, elem2)
	//	slice = append(slice, anotherSlice...)
	// As a special case, it is legal to append a string to a byte slice, like this:
	//	slice = append([]byte("hello "), "world"...)


	var a []int
	// 需要一个接收方，因为append在扩容时，会返回新对象
	b := append(a,1)
	t.Log(b)
	t.Log(a)

	// 原来的对象保持不变
	c := make([]int, 4)
	d := append(c,4)
	t.Log(c)
	t.Log(d)
}
