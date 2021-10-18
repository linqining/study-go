package _11018

//递归和栈操作逆序栈
type Stack struct{
	i int
	data [10]int
}

func (s *Stack)Push(k int){
	s.data[s.i] = k
	s.i++
}

func(s *Stack) Pop()(ret int){
	s.i--
	ret = s.data[s.i]
	return
}

func(s *Stack) Peek()(ret int){
	ret = s.data[s.i-1]
	return
}

func (s *Stack) IsEmpty() bool{
	return s.i ==0
}

func (s *Stack) getAndRemoveLastElement()int{
	elm := s.Pop()
	if s.IsEmpty(){
		return elm;
	}else{
		last := s.getAndRemoveLastElement()
		s.Push(elm)
		return last
	}
}

func (s *Stack) Reverse(){
	if s.IsEmpty(){
		return;
	}
	elm := s.getAndRemoveLastElement()
	s.Reverse()
	s.Push(elm)
}

// 实现栈的排序
// 允许申请一个新的栈，不允许使用其他数据结构
func (s *Stack) Sort(){
	nStack := Stack{}
	elmCount :=s.i
	for i:=elmCount;i>0;i--{
		minCount :=0
		minElm := s.Peek()
		for j:=0;j<i;j++{
			tmpElm := s.Pop()
			if tmpElm < minElm{
				minElm = tmpElm
			}
			nStack.Push(tmpElm)
		}
		s.Push(minElm)
		for !nStack.IsEmpty(){
			popElm := nStack.Pop()
			if popElm == minElm && minCount ==0{
				minCount ++
				continue
			}
			s.Push(popElm)
		}
	}
}