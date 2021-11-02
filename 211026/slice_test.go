package _11026

import (
	"fmt"
	"testing"
)

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

type tst1 struct{
	a int
	b []int
}

type tst2 struct{
	a int
	b int
}

// 包含引用类型的不能比较
func TestStructCompare(t *testing.T){
	//o1 := tst1{a:0,b:[]int{1}}
	//o2 := tst1{a:0, b:[]int{1}}
	//if o1==o2{
	//	t.Log("o1等于o2")
	//}else{
	//	t.Log("o1不等于o2")
	//}

	o3 := tst2{a:1,b:2}
	o4 := tst2{a:1,b:2}

	if o3==o4{
		t.Log("o3等于o4")
	}else{
		t.Log("o3不等于o4")
	}
}

func TestPassSlice(t *testing.T) {
	var arr = make([]int,0,10)
	doSomeHappyThings(arr)
	fmt.Println(arr, len(arr), cap(arr), "after return")
}

func doSomeHappyThings(arr []int) {
	arr = append(arr, 1)
	fmt.Println(arr, "after append")
}
