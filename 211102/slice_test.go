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

func TestMaxSubArr2(t *testing.T) {
	res := MaxSubArr2([]int{-1,2,1,6,0,-1,9,4,3})
	t.Log(res)
}

func TestMaxdiff(t *testing.T) {
	res := Maxdiff([]int{5,1})
	t.Log(res)
	res = Maxdiff([]int{5,6})
	t.Log(res)
}

func TestGrayCode(t *testing.T) {
	res := GrayCode(1)
	t.Log(res)
	res = GrayCode(2)
	t.Log(res)
	res = GrayCode(3)
	t.Log(res)
}

func TestGetWonPrice(t *testing.T) {
	wonmon:=GetWonPrice([]int{1,3,1,1,4})
	t.Log(wonmon)

	wonmon=GetWonPrice([]int{1,3,2,1,4})
	t.Log(wonmon)

	wonmon=GetWonPrice([]int{4,1,2})
	t.Log(wonmon)
}


func TestBinarySearch(t *testing.T) {
	idx :=BinarySearch([]int{1,3,5,7,9},3)
	t.Log(idx)
	idx =BinarySearch([]int{1,3,5,7,9},0)
	t.Log(idx)
	idx =BinarySearch([]int{1,3,5,7,9},2)
	t.Log(idx)
	idx =BinarySearch([]int{1,3,5,7},2)
	t.Log(idx)
}