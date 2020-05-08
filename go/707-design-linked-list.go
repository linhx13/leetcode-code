package main

type Node struct {
	val  int
	next *Node
}

type MyLinkedList struct {
	head *Node
	tail *Node
	size int
}

/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{
		head: nil,
		tail: nil,
		size: 0,
	}
}

/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.size {
		return -1
	}
	node := this.head
	for i := 0; i < index; i++ {
		node = node.next
	}
	return node.val
}

/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int) {
	this.head = &Node{val: val, next: this.head}
	if this.size == 0 {
		this.tail = this.head
	}
	this.size++
}

/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int) {
	node := &Node{val: val}
	if this.size == 0 {
		this.head = node
		this.tail = node
	} else {
		this.tail.next = node
		this.tail = node
	}
	this.size++
}

/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int) {
	if index < 0 || index > this.size {
		return
	}
	if index == 0 {
		this.AddAtHead(val)
	} else if index == this.size {
		this.AddAtTail(val)
	} else {
		prev := &Node{0, this.head}
		for i := 0; i < index; i++ {
			prev = prev.next
		}
		prev.next = &Node{val: val, next: prev.next}
		this.size++
	}
}

/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int) {
	if index < 0 || index >= this.size {
		return
	}
	prev := &Node{0, this.head}
	for i := 0; i < index; i++ {
		prev = prev.next
	}
	prev.next = prev.next.next
	if index == 0 {
		this.head = prev.next
	} else if index == this.size-1 {
		this.tail = prev
	}
	this.size--
}

func main() {

}
