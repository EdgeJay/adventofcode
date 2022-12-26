package data

import (
	"fmt"
	"strings"
)

type LinkedListNode[V comparable] struct {
	Value  V
	Parent *LinkedListNode[V]
}

type LinkedList[V comparable] struct {
	list []*LinkedListNode[V]
}

func (l *LinkedList[V]) AddNode(value V) {
	var parent *LinkedListNode[V] = nil
	if len(l.list) > 0 {
		parent = l.list[len(l.list)-1]
	}

	node := &LinkedListNode[V]{
		Value:  value,
		Parent: parent,
	}
	l.list = append(l.list, node)
}

func (l *LinkedList[V]) StartsAndEndWithValues(start, end V) bool {
	return l.list[0].Value == start && l.list[len(l.list)-1].Value == end
}

func (l *LinkedList[V]) Duplicate() *LinkedList[V] {
	listCopy := make([]*LinkedListNode[V], len(l.list))
	copy(listCopy, l.list)
	return &LinkedList[V]{
		list: listCopy,
	}
}

func (l *LinkedList[V]) Size() int {
	return len(l.list)
}

func (l *LinkedList[V]) GetLastNode() *LinkedListNode[V] {
	return l.list[len(l.list)-1]
}

func (l *LinkedList[V]) PrintNodes() string {
	arr := make([]string, len(l.list))
	for idx, node := range l.list {
		arr[idx] = fmt.Sprint(node.Value)
	}
	return strings.Join(arr, ", ")
}

func NewLinkedList[V comparable]() *LinkedList[V] {
	return &LinkedList[V]{
		list: make([]*LinkedListNode[V], 0),
	}
}
