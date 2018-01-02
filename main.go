package main

import (
	"fmt"
	"go-algorithm/algorithm/link"
)

func main() {
	n1 := &link.Node{}
	n2 := &link.Node{
		121,
		"23423",
		nil,
	}
	n3 := &link.Node{
		1211,
		324,
		nil,
	}
	n4 := &link.Node{
		12111,
		12,
		nil,
	}

	n1.Next = n2
	n2.Next = n3
	n3.Next = n4

	n := n1.Frist()
	fmt.Println("frist data：", n.Data)
	fmt.Println("frsit T：", n.T)

	nl := n1.Last()
	fmt.Println("last data：",nl.Data)

	fmt.Println("length:",n1.Length())

	rn,err := n1.Get(3)
	fmt.Println("get: ",rn,err)
}