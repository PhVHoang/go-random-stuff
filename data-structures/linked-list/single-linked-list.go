package main

type Node struct {
	val  int
	next *Node
}

type SingleLinkedList struct {
	head *Node
}

// to avoid mistakes when using pointer vs struct for new node creation
func newNode(val int) *Node {
	return &Node{val, nil}
}

func (ll *SingleLinkedList) addAtBegin(val int) {
	n := newNode(val)
	n.next = ll.head
	ll.head = n
}

func (ll *SingleLinkedList) addAtEnd(val int) {
	n := newNode(val)
	if ll.head == nil {
		ll.head = n
		return
	}
	cur := ll.head
	for ; cur.next != nill; cur = cur.next {
	}
	cur.next = n

}
func (ll *SingleLinkedList) deleteAtBegin() {
	if ll.head == nill {
		return
	}

}
