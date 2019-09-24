package doublelinkedlist

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestListItem_Prev(t *testing.T) {
	current := &ListItem{value: 8}
	previous := &ListItem{next: current, value: 7}
	current.prev = previous

	got := current.Prev()
	require.Equal(t, previous, got)
}

func TestListItem_Value(t *testing.T) {
	item := &ListItem{value: "hello, друзья"}
	got := item.Value()
	require.Equal(t, "hello, друзья", got)
}

func TestListItem_Next(t *testing.T) {
	current := &ListItem{value: []int{1, 2, 3}}
	next := &ListItem{value: []int{4, 5, 6}, prev: current}
	current.next = next

	got := current.Next()
	require.Equal(t, next, got)
}

func TestList_Len(t *testing.T) {
	first := &ListItem{value: 1}
	mid := &ListItem{value: 2}
	last := &ListItem{value: 3}

	first.next = mid
	mid.prev = first
	mid.next = last
	last.prev = mid

	list := NewList(first, last)
	require.Equal(t, 3, list.Len())
}

func TestList_LenForEmptyList(t *testing.T) {
	list := NewList(nil, nil)
	require.Equal(t, 0, list.Len())
}

func TestList_LenForSingleItemList(t *testing.T) {
	first := &ListItem{value: 1}
	list := NewList(first, nil)
	require.Equal(t, 1, list.Len())
}

func TestList_First(t *testing.T) {
	first := &ListItem{value: 1}
	mid := &ListItem{value: 2}
	last := &ListItem{value: 3}

	first.next = mid
	mid.prev = first
	mid.next = last
	last.prev = mid

	list := NewList(first, last)
	require.Equal(t, first, list.First())
}

func TestList_FirstForEmptyList(t *testing.T) {
	list := NewList(nil, nil)
	require.Nil(t, list.First())
}

func TestList_FirstForSingleItemList(t *testing.T) {
	first := &ListItem{value: 1}
	list := NewList(first, nil)
	require.Equal(t, first, list.First())
}

func TestList_Last(t *testing.T) {
	first := &ListItem{value: 1}
	mid := &ListItem{value: 2}
	last := &ListItem{value: 3}

	first.next = mid
	mid.prev = first
	mid.next = last
	last.prev = mid

	list := NewList(first, last)
	require.Equal(t, last, list.Last())
}

func TestList_LastForEmptyList(t *testing.T) {
	list := NewList(nil, nil)
	require.Nil(t, list.Last())
}

func TestList_LastForSingleItemList(t *testing.T) {
	first := &ListItem{value: 1}
	list := NewList(first, nil)
	require.Equal(t, first, list.Last())
}

func TestList_RemoveMiddleItem(t *testing.T) {
	first := &ListItem{value: 1}
	mid := &ListItem{value: 2}
	last := &ListItem{value: 3}

	first.next = mid
	mid.prev = first
	mid.next = last
	last.prev = mid

	list := NewList(first, last)

	err := list.Remove(mid)

	require.Nil(t, err)
	require.Equal(t, list.Len(), 2)
	require.Equal(t, list.First(), first)
	require.Equal(t, list.Last(), last)
}

func TestList_RemoveFirstItem(t *testing.T) {
	first := &ListItem{value: 1}
	mid := &ListItem{value: 2}
	last := &ListItem{value: 3}

	first.next = mid
	mid.prev = first
	mid.next = last
	last.prev = mid

	list := NewList(first, last)

	err := list.Remove(first)

	require.Nil(t, err)
	require.Equal(t, 2, list.Len())
	require.Equal(t, mid, list.First())
	require.Equal(t, last, list.Last())
}

func TestList_RemoveLastItem(t *testing.T) {
	first := &ListItem{value: 1}
	mid := &ListItem{value: 2}
	last := &ListItem{value: 3}

	first.next = mid
	mid.prev = first
	mid.next = last
	last.prev = mid

	list := NewList(first, last)

	err := list.Remove(last)

	require.Nil(t, err)
	require.Equal(t, 2, list.Len())
	require.Equal(t, first, list.First())
	require.Equal(t, mid, list.Last())
}

func TestList_RemovePassNotExistingListItem(t *testing.T) {
	first := &ListItem{value: 1}
	mid := &ListItem{value: 2}
	last := &ListItem{value: 3}

	first.next = mid
	mid.prev = first
	mid.next = last
	last.prev = mid

	list := NewList(first, last)

	err := list.Remove(&ListItem{value: "i am not existing in list"})
	require.EqualError(t, err, "passed item not found in list")
}

func TestList_RemoveTryForEmptyList(t *testing.T) {
	list := NewList(nil, nil)
	err := list.Remove(&ListItem{value: "i am not existing in list"})
	require.EqualError(t, err, "passed item not found in list")
}

func TestList_RemoveSingleItemList(t *testing.T) {
	first := &ListItem{value: 1}
	list := NewList(first, nil)
	err := list.Remove(first)

	require.Nil(t, err)
	require.Equal(t, 0, list.Len())
}

func TestList_PushBack(t *testing.T) {
	first := &ListItem{value: 1}
	list := NewList(first, nil)
	list.PushBack("new item")

	require.Equal(t, 2, list.Len())
	require.Equal(t, "new item", list.Last().Value())
}

func TestList_PushBackForEmptyList(t *testing.T) {
	list := NewList(nil, nil)
	list.PushBack("new item")

	require.Equal(t, 1, list.Len())
	require.Equal(t, "new item", list.Last().Value())
}

func TestList_PushFront(t *testing.T) {
	first := &ListItem{value: 1}
	list := NewList(first, nil)
	list.PushFront("new item")

	require.Equal(t, 2, list.Len())
	require.Equal(t, "new item", list.First().Value())
}

func TestList_PushFrontForEmptyList(t *testing.T) {
	list := NewList(nil, nil)
	list.PushFront("new item")

	require.Equal(t, 1, list.Len())
	require.Equal(t, "new item", list.First().Value())
}
