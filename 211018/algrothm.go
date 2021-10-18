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