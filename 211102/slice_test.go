package _11102

import "testing"

func TestInitSlice(t *testing.T) {
	InitSlice()
}


func TestChanPanic(t *testing.T) {
	ChanPanic()
}

func TestPrintNums(t *testing.T) {
	PrintNums()
}

func TestFindMiss(t *testing.T) {
	arr := []int{3,3,1}
	count := FindMiss(arr)
	t.Log(count)
}

func TestGetNDan(t *testing.T) {
	t.Log(GetNDan(1))
	t.Log(GetNDan(2))
	t.Log(GetNDan(3))
	t.Log(GetNDan(10))
	t.Log(GetNDan(1000000000))
}

func TestMaxSubArr(t *testing.T) {
	_,_,res := MaxSubArr([]int{-1,2,1,6,0,-1,9,4,3},false)
	t.Log(res)
}