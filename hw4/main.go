/*
Пакет, содержаший основы работы с двусвязаным списком.
*/
package doublelinkedlist

import "fmt"

// Элемент двусвязного списка.
type ListItem struct {
	prev  *ListItem
	next  *ListItem
	value interface{}
}

func (i *ListItem) Next() *ListItem {
	return i.next
}

func (i *ListItem) Prev() *ListItem {
	return i.prev
}

func (i *ListItem) Value() interface{} {
	return i.value
}

// Двусвязный список.
type List struct {
	first *ListItem
	last  *ListItem
}

func (l *List) Len() int {
	s := 0
	current := l.first
	for current != nil {
		current = current.next
		s += 1
	}
	return s
}

func (l *List) First() *ListItem {
	return l.first
}

func (l *List) Last() *ListItem {
	if l.last != nil {
		return l.last
	} else {
		return l.first
	}
}

func (l *List) Remove(item *ListItem) error {
	current := l.first
	for current != item && current != nil {
		current = current.next
	}
	if current == nil {
		return fmt.Errorf("passed item not found in list")
	}

	prev := current.prev
	next := current.next

	if prev != nil {
		prev.next = next
	} else {
		l.first = next
	}
	if next != nil {
		next.prev = prev
	} else {
		l.last = prev
	}

	return nil
}

func (l *List) PushBack(v interface{}) {
	currentLast := l.Last()
	item := &ListItem{prev: currentLast, value: v}
	if currentLast != nil {
		currentLast.next = item
		l.last = item
	} else {
		l.first = item
	}
}

func (l *List) PushFront(v interface{}) {
	currentFirst := l.First()
	item := &ListItem{next: currentFirst, value: v}
	if currentFirst != nil {
		currentFirst.prev = item
	}
	l.first = item
}

func NewList(first *ListItem, last *ListItem) *List {
	return &List{
		first: first,
		last:  last,
	}
}
