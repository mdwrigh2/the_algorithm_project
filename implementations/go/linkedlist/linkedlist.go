// linkedlist.go - linked list functions/algorithms

package linkedlist

// DATA
// ----
type Data int

func (d *Data) eq(a *Data) bool {
	if a == nil || d == nil {
		return false
	}
	return *d == *a
}

// helper
func NewData(i Data) *Data {
	d := new(Data)
	*d = i
	return d
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
	if *head != nil { 
		(*head).prev = item // only for doubly-linked lists
	}
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
	item.prev = previous // only for doubly-linked lists
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
	item.prev = *tail // only for doubly-linked lists
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
		if *head != nil {
			(*head).prev = nil // only for doubly-linked lists
		}
		return true
	}
	current := *head
	for current != nil {
		if current.next == delete {
			current.next = delete.next
			if current.next != nil {
				current.next.prev = current // only for doubly-linked lists
			}
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

// This is more useful in C, where every pointer isnt initialized to nil
func NewStack(stack **Node) bool {
	if stack == nil { // bogus pointer, can't do
		return false
	}
	*stack = nil
	return true
}

func Push(stack **Node, d *Data) bool {
	if stack == nil { // bogus pointer, can't do
		return false
	}
	item := NewNode(d);
	if *stack == nil {
		*stack = item
		return true
	}
	item.next = *stack
	(*stack).prev = item // only for doubly-linked list
	*stack = item
	return true
}

// in C a return of nil might mean failure, however in go
// we are accepting nil as a valid *Data for our list
func Pop(stack **Node) (d *Data, ok bool) {
	if stack == nil { // bogus pointer, can't do
		return nil,false
	}
	if *stack == nil { // empty list, can't pop
		return nil,false
	}
	item := *stack
	*stack = item.next
	if *stack != nil {
		(*stack).prev = nil // only for doubly-linked list
	}
	data := item.data
	// free(item) but we're garbage-collected so not needed
	return data, true
}

// This function is also pretty useless in Go, but useful in C
func DeleteStack(stack **Node) bool {
	if stack == nil { // bogus stack, can't do
		return false
	}
	for *stack != nil {
		next := (*stack).next
		if next != nil {
			next.prev = nil // only for doubly-linked list
		}
		// free(current) but we're garbage-collected so not needed
		*stack = next
	}
	return true
}


