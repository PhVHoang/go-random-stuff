package main

import "fmt"

type Node struct {
  val int
  next *Node
  prev *Node
}

type DoubleLinkedList {
  head *Node
}

func newNode(val int) *Node {
  node := &Node{}
  node.val = val
  node.next = nil
  node.prev = nil
  return node
}

func (ll * DoubleLinkedList) addAtBegin(val int) {
  n := newNode(val)
  n.next = ll.head
  ll.head = n
}

func (ll *DoubleLinkedList) addAtEnd(val int) {
  n := newNode(val)
  if ll.head == nil {
    ll.head = n
    return
  }

  cur := ll.head
  for ; cur.next != nil; cur = cur.next {
  }
  cur.next = n
  n.prev = cur
}

func (ll *DoubleLinkedList) deleteAtBegin() int{
  if ll.head == nil {
    return -1
  }
  cur := ll.head
  ll.head = cur.next
  if ll.head != nil {
    ll.head.prev = nil
  }
  return cur.val
}

func (ll *DoubleLinkedList) deleteAtEnd() int{
  if ll.head == nil {
    return -1
  }

  if ll.head.next != nil {
    ll.deleteAtBegin()
  }
  cur := ll.head
  for ; cur.next.next != nil; cur = cur.next{
  }
  retval := cur.next.val
  cur.next = nil
  return retval
}

func (ll *DoubleLinkedList) count() int {
  var ctr int = 0
  for cur := ll.head; cur.next != nil; cur = cur.next{
    ctr += 1
  }
  return ctr
}

func (ll *DoubleLinkedList) reverse() {
  var prev, next *Node
  cur := ll.head
  for cur != nil {
    next = cur.next
    cur.next = prev
    cur.prev = next
    prev = cur
    cur = next
  }
  ll.head = prev
}

func (ll *DoubleLinkedList) displayReverse() {
  if ll.head == nil {
    return
  }
  var cur *Node
  for cur = ll.head; cur.next != nil; cur = cur.next {
  }
  for ; cur != nil; cur = cur.prev {
    fmt.Println(cur.val, " ")
  }
  fmt.Println("\n")
}


func main() {

}
