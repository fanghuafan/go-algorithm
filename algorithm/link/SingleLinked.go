package link

import "errors"

// data struct
type Node struct {
	Data int         // data
	T    interface{} // test
	Next *Node       // next node
}

// get the frist node
func (t *Node) Frist() *Node {
	if t.Next == nil {
		return nil;
	}
	return t.Next
}

// get the last node
func (t *Node) Last() *Node {
	if t.Next == nil {
		return nil;
	}

	temp := t
	for temp.Next != nil {
		temp = temp.Next
		if temp.Next == nil {
			return temp
		}
	}
	return nil
}

// get the len
func (t *Node) Length() (count int64) {
	if t.Next == nil {
		return 0
	}
	for t.Next != nil {
		count ++
		t = t.Next
	}
	return
}

// insert node
func (t *Node) Insert(new *Node, index int) (bool, error) {
	if t == nil {
		return false, errors.New("node nil")
	}
	count := 0
	for t.Next != nil {
		count ++
		if count == index {
			new.Next = t.Next
			t.Next = new
			return true, nil
		}
		t = t.Next
	}
	return false, errors.New("insert index no found")
}

// delete node
func (t *Node) Delete(index int) (bool, error) {
	if t == nil {
		return false, errors.New("node nil")
	}
	count := 0
	for t.Next != nil {
		count ++
		if count == index {
			if t.Next.Next != nil {
				t.Next = t.Next.Next
				return true, nil
			}
		}
		t = t.Next
	}
	return false, errors.New("delete target no found")
}

// delete node
func (t *Node) DeleteByNode(target *Node) (bool, error) {
	if t == nil {
		return false, errors.New("node nil")
	}
	for t.Next != nil {
		if t.Data == target.Data {
			if t.Next.Next != nil {
				t.Next = t.Next.Next
				return true, nil
			}
		}
		t = t.Next
	}
	return false, errors.New("delete target no found")
}

// get the node by index
func (t *Node) Get(index int) (*Node, error) {
	if t == nil {
		return nil, errors.New("node nil")
	}
	count := 1
	for t.Next != nil {
		if count == index {
			return t, nil
		}
		count ++
		t = t.Next
	}
	return nil, errors.New("delete target index no found")
}
