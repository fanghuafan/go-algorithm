package _select

import (
	"testing"
)

func TestBubble(t *testing.T) {
	var arr = []int{1,4,5,7,2,3,10,-1,-12,23,324,23423,-1212}
	SelectSort(arr)
	t.Log(arr)
	//i := 4
	//j := 1
	//
	//fmt.Println("----:", i-j)
}
