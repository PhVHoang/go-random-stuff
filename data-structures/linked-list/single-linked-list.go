package main

import "fmt"

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
	for ; cur.next != nil; cur = cur.next {
	}
	cur.next = n

}

func (ll *SingleLinkedList) deleteAtBegin() int {
	if ll.head == nil {
		return -1
	}
	cur := ll.head
	ll.head = cur.next
	return cur.val
}

func (ll *SingleLinkedList) deleteAtEnd() int {
	if ll.head == nil {
		return -1
	}

	cur := ll.head
	for ; cur.next.next != nil; cur = cur.next {
	}
	retval := cur.next.val
	cur.next = nil
	return retval
}

func (ll *SingleLinkedList) count() int {
	var ctr int = 0
	cur := ll.head
	for ; cur.next != nil; cur = cur.next {
		ctr += 1
	}
	return ctr
}

func (ll *SingleLinkedList) reverse() {
	var prev, next *Node
	cur := ll.head
	for cur != nil {
		next = cur.next
		cur.next = prev
		prev = cur
		cur = next
	}

	ll.head = prev
}

func (ll *SingleLinkedList) display() {
	for cur := ll.head; cur != nil; cur = cur.next {
		fmt.Print(cur.val, " ")
	}
}

func main() {
	ll := SingleLinkedList{}

	ll.addAtBegin(10)
	ll.addAtEnd(20)
	ll.display()
	fmt.Println()
	ll.addAtBegin(30)
	ll.display()
	fmt.Println()
	ll.reverse()
	ll.display()
	fmt.Println()

	fmt.Print(ll.deleteAtBegin(), "\n")
	ll.display()

	fmt.Print(ll.deleteAtEnd(), "\n")
	ll.display()

}
