package bubble

import (
	"testing"
)

func TestBubble(t *testing.T) {
	var arr = []int{1,4,5,7,2,3,10,-1,-12}
	num := SortUpgrade(arr)
	t.Log(num)
}