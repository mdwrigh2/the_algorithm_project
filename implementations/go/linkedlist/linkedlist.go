// linkedlist.go - linked list functions/algorithms

package linkedlist

// DATA
// ----
type Data int

func (d *Data) eq(a *Data) bool {
	return *d == *a
}

// SINGLY LINKED LIST
// ------------------
type Node struct {
	next  *Node
	prev  *Node
	child *Node
	data  *Data
}

func NewNode(d *Data) *Node {
	item := new(Node)
	item.data = d
	return item
}

func NewNodeData(d Data) *Node {
	data := new(Data)
	item := new(Node)
	*data = d
	item.data = data
	return item
}

func Prepend(head **Node, d *Data) bool {
	if head == nil { // bogus head, no list
		return false
	}
	item := NewNode(d)
	item.next = *head // handles empty or full lists
	*head = item
	return true
}

func Append(head **Node, d *Data) bool {
	if head == nil { // bogus head, no list
		return false
	}
	item := NewNode(d)
	current := *head
	var previous *Node
	for current != nil { // traverse to end
		previous = current
		current = current.next
	}
	if previous == nil { // empty list
		*head = item
		return true
	}
	previous.next = item
	return true
}

func AppendTail(head **Node, tail **Node, d *Data) bool {
	// NOTE: Assumes head and tail are consistent
	if head == nil || tail == nil { // bogus list
		return false
	}
	item := NewNode(d)
	if *tail == nil { // empty list
		*tail = item
		*head = item
		return true
	}
	(*tail).next = item
	*tail = item
	return true
}

func Find(head *Node, d *Data) *Node {
	current := head
	for current != nil {
		if d.eq(current.data) {
			return current
		}
		current = current.next
	}
	return nil
}

func Delete(head **Node, delete *Node) bool {
	if head == nil { // bogus head, no list
		return false
	}
	if delete == nil { // bogus delete, you can always remove nil
		return true
	}
	if *head == delete {
		*head = delete.next
		return true
	}
	current := *head
	for current != nil {
		if current.next == delete {
			current.next = delete.next
			// free(delete) but we're garbage-collected so not needed
			return true
		}
		current = current.next
	}
	return false // didn't find it
}

// Bit of a useless function, but in C this is necessary
func DeleteList(head **Node) bool {
	if head == nil { // bogus head, no list
		return false
	}
	current := *head
	for current != nil {
		next := current.next
		// free(current) but we're garbage-collected so not needed
		current = next
	}
	*head = nil // equivalent to this whole function
	return true
}

