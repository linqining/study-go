package _10917

import "fmt"

type Slice []int


func NewSlice() Slice{
	return make(Slice,0)
}

func (s *Slice) Add(elem int) *Slice{
	*s = append(*s, elem)
	fmt.Print(elem)
	return s
}

// add(1) 会优先add(3)执行
// 同理add(4) 会优先add(3)
// add(2) 压入栈中等待return执行
func SliceTest(){
	s:= NewSlice()
	defer s.Add(1).Add(4).Add(2)
	s.Add(3)
}
