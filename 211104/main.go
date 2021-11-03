package _11104

import "fmt"

type N int

func (n N)test(){
	fmt.Println(n)
}

func CaptureVar(){
	var n N =10
	p := &n

	n++
	f1:=n.test
	n++

	f2:=p.test
	n++
	fmt.Println(n)
	f1()
	f2()
}


func CompareInterface(){
	var x interface{}
	var y interface{} = []int{2,3}
	_=x==x
	_=x==y
	_=y==y
}