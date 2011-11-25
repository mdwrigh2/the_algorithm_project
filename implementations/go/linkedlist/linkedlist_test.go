package linkedlist

import "testing"

type dataPair struct {
	a, b Data
	eq   bool
}

var dataPairs = []dataPair{
	dataPair{Data(1), Data(1), true},
	dataPair{Data(1), Data(2), false},
	dataPair{Data(0), Data(0), true},
	dataPair{Data(8), Data(0), false},
}

// just for testing's sake, yes I know it's pointless
func TestDataEq(t *testing.T) {
	for _, dp := range dataPairs {
		res := dp.a.eq(&dp.b)
		if dp.eq != res {
			t.Errorf("Data.eq() %v == %v got %t, expected %t", dp.a, dp.b, res, dp.eq)
		}
	}
}

func TestPrepend(t *testing.T) {
	// test failure case (no list at all, bogus head ptr)
	if Prepend(nil, nil) {
		t.Errorf("Prepend(nil,nil) got true, expected false")
	}
	// test empty list case (*head == nil)
	var head *Node
	Prepend(&head, nil)
	if head == nil {
		t.Errorf("Prepend() failed on empty list")
	}
	// test on actual list
	head = new(Node)
	old := head
	Prepend(&head, nil)
	if head == nil || head.next != old || head.next.next != nil || head.next.prev != head {
		t.Errorf("Prepend() failed on list")
	}
}

func TestAppend(t *testing.T) {
	// test failure case (no list at all, bogus head ptr)
	if Append(nil, nil) {
		t.Errorf("Append(nil,nil) got true, expected false")
	}
	// test empty list case (*head == nil)
	var head *Node
	Append(&head, nil)
	if head == nil {
		t.Errorf("Append() failed on empty list")
	}
	// test on actual list
	head = new(Node)
	old := head
	Append(&head, nil)
	if head != old || head.next == nil || head.next.next != nil || head.next.prev != head {
		t.Errorf("Append() failed on list")
	}
}

func TestAppendTail(t *testing.T) {
	// test failure case (no list at all, bogus head ptr)
	if AppendTail(nil, nil, nil) {
		t.Errorf("AppendTail(nil,nil) got true, expected false")
	}
	// test empty list case (*head == nil)
	var head, tail *Node
	AppendTail(&head, &tail, nil)
	if head == nil || head != tail {
		t.Errorf("AppendTail() failed on empty list")
	}
	// test on actual list
	head = new(Node)
	tail = head
	old := head
	AppendTail(&head, &tail, nil)
	if head != old || head == tail || head.next != tail || tail.next != nil || tail.prev != head {
		t.Errorf("AppendTail() failed on list")
	}
}

func TestFind(t *testing.T) {
	// test empty list case
	if found := Find(nil,nil) ; found != nil {
		t.Errorf("Find(nil,nil) got a pointer, expected nil")
	}
	// test one-item list
	head := NewNodeData(1)
	if found := Find(head,NewData(1)) ; found != head {
		t.Errorf("Find() failed on one-item list")
	}
	// test two-item list
	head = NewNodeData(1)
	tail := NewNodeData(2)
	head.next = tail
	if found := Find(head,NewData(2)) ; found != tail {
		t.Errorf("Find() failed on two-item list")
	}
}

func TestDelete(t *testing.T) {
	// test bogus list case
	if Delete(nil,nil) {
		t.Errorf("Delete(nil,nil) got true, expected false")
	}
	// test bogus delete case
	var head *Node
	if !Delete(&head,nil) {
		t.Errorf("Delete(&head,nil) got false, expected true")
	}
	// test empty list
	tail := new(Node)
	if Delete(&head,tail) {
		t.Errorf("Delete() succeeded on empty list")
	}
	// test delete head (one-item list) case
	head = new(Node)
	if !Delete(&head,head) || head != nil {
		t.Errorf("Delete(&head,head) failed on one-item list")
	}
	// test delete not in list
	head = new(Node)
	tail = new(Node)
	old := head
	if Delete(&head,tail) || head != old {
		t.Errorf("Delete(&head,tail) failed on delete not in list")
	}
	// test delete head (two-item list) case
	head = new(Node)
	tail = new(Node)
	old = head
	head.next = tail
	tail.prev = head
	if !Delete(&head,head) || head != tail || head.next != nil || head.prev != nil {
		t.Errorf("Delete(&head,head) failed on two-item list")
	}
	// test delete tail (two-item list) case
	head = new(Node)
	tail = new(Node)
	old = head
	head.next = tail
	if !Delete(&head,tail) || head != old || head.next != nil {
		t.Errorf("Delete(&head,tail) failed on two-item list")
	}
}

func TestDeleteList(t *testing.T) {
	// test bogus list case
	if DeleteList(nil) {
		t.Errorf("DeleteList(nil) got true, expected false")
	}
	// test empty list
	var head *Node
	if !DeleteList(&head) || head != nil {
		t.Errorf("DeleteList() failed on empty list")
	}
	// test delete head (one-item list) case
	head = new(Node)
	if !DeleteList(&head) || head != nil {
		t.Errorf("DeleteList() failed on one-item list")
	}
	// test delete tail (two-item list) case
	head = new(Node)
	tail := new(Node)
	head.next = tail
	if !DeleteList(&head) || head != nil {
		t.Errorf("DeleteList() failed on two-item list")
	}
}

func TestNewStack(t *testing.T) {
	// test bogus pointer case
	if NewStack(nil) {
		t.Errorf("NewStack(nil) got true, expected false")
	}
	// test empty list
	head := new(Node)
	if !NewStack(&head) || head != nil {
		t.Errorf("NewStack() failed on existing list")
	}
}

func TestPush(t *testing.T) {
	// test failure case (no list at all, bogus head ptr)
	if Push(nil, nil) {
		t.Errorf("Prepend(nil,nil) got true, expected false")
	}
	// test empty list case (*head == nil)
	var head *Node
	Push(&head, nil)
	if head == nil {
		t.Errorf("Push() failed on empty list")
	}
	// test on actual list
	head = new(Node)
	old := head
	Prepend(&head, nil)
	if head == nil || head.next != old || head.next.next != nil || head.next.prev != head {
		t.Errorf("Push() failed on list")
	}
}

func TestPop(t *testing.T) {
	// test bogus list case
	if _, ok := Pop(nil) ; ok {
		t.Errorf("Pop(nil) got true, expected false")
	}
	// test empty list case
	var head *Node
	if _, ok := Pop(&head) ; ok || head != nil {
		t.Errorf("Pop() succeeded on empty list")
	}
	// test one-item list
	head = NewNodeData(1)
	if d, ok := Pop(&head) ; !ok || !NewData(1).eq(d) || head != nil {
		t.Errorf("Pop() failed on one-item list")
	}
	// test two-item list
	head = NewNodeData(1)
	tail := NewNodeData(2)
	head.next = tail
	tail.prev = head
	if d, ok := Pop(&head) ; !ok || !NewData(1).eq(d) || head != tail || head.prev != nil {
		t.Errorf("Pop() failed on two-item list")
	}
}

func TestDeleteStack(t *testing.T) {
	// test bogus list case
	if DeleteStack(nil) {
		t.Errorf("DeleteStack(nil) got true, expected false")
	}
	// test empty list case
	var head *Node
	if !DeleteStack(&head) {
		t.Errorf("DeleteStack() failed on empty list")
	}
	// test one-item list
	head = NewNodeData(1)
	if !DeleteStack(&head) || head != nil {
		t.Errorf("DeleteStack() failed on one-item list")
	}
	// test two-item list
	head = NewNodeData(1)
	tail := NewNodeData(2)
	head.next = tail
	tail.prev = head
	if !DeleteStack(&head) || head != nil {
		t.Errorf("DeleteStack() failed on two-item list")
	}
}


