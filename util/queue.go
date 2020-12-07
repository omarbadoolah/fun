package util

import (
	"fmt"
)

// Node - basic element of queue, list, etc
type Node struct {
	value interface{}
	next  *Node
}

// Queue - data structure that adds elements to the end and removes from the start
type Queue struct {
	head *Node
	tail *Node
}

// Add - add an element to the queue
func (q *Queue) Add(value interface{}) error {
	newNode := Node{value: value, next: nil}
	if q.tail != nil {
		q.tail.next = &newNode
		q.tail = q.tail.next
	} else {
		if q.head != nil {
			return fmt.Errorf("Add(): Tail is empty, but head isn't")
		}
		q.head = &newNode
		q.tail = q.head
	}
	return nil
}

// Pop - remove and return the first element of the queue
// returns nil if empty
func (q *Queue) Pop() *Node {
	poppedNode := q.head
	if q.head != nil {
		q.head = q.head.next
	}
	return poppedNode
}

// List - returns current contents of queue as a list
func (q *Queue) List() []interface{} {
	list := make([]interface{}, 0)
	current := q.head
	for current != nil {
		list = append(list, current.value)
		current = current.next
	}
	return list
}
