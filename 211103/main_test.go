package _11103

import "testing"

func TestGenBoard(t *testing.T) {
	GenBoard()
}

func TestGetMostBonus(t *testing.T) {
	GetMostBonus()
}

func TestJumpToX(t *testing.T) {
	step := JumpToX(999)
	t.Log(step)
}

func TestFindFriends(t *testing.T) {
	FindFriends(10)
}

func TestCountBitDiff(t *testing.T) {
	diffCount := CountBitDiff(1999,2299)
	t.Log(diffCount)
}


func TestRestaruantMaxIncome(t *testing.T) {
	res :=RestaruantMaxIncome([]int{2,4,2},[][]int{[]int{1,3},[]int{3,5},[]int{3,7},[]int{5,9},[]int{1,10}})
	t.Log(res)
}

func TestSerialMax(t *testing.T) {
	data := []QScore{{1,10},{1,5},{2,4},{3,9},{4,8}}
	score,_:=SerialMax(data)
	t.Log(score)
}

func TestSumTree(t *testing.T) {
	SumTree()
}