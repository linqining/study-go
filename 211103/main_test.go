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