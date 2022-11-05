package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLen(t *testing.T) {
	list := List{}
	assert.Equal(t, 0, list.Len(), "Uninitialized list")

	list.PushFront(1)
	assert.Equal(t, 1, list.Len(), "List with one item")

	list.PushFront(2)
	assert.Equal(t, 2, list.Len(), "List with more than one item")
}

func TestFirstItem(t *testing.T) {
	a := Item{}
	b := Item{prev: &a}
	a.next = &b
	list := List{
		first: &a,
		last:  &b,
	}

	assert.Equal(t, &a, list.First())
}

func TestLastItem(t *testing.T) {
	a := Item{}
	b := Item{prev: &a}
	a.next = &b
	list := List{
		first: &a,
		last:  &b,
	}

	assert.Equal(t, &b, list.Last())
}

func TestPushFrontToEmptyList(t *testing.T) {
	list := List{}

	list.PushFront(1)

	a := list.First()
	b := list.Last()

	assert.Equal(t, 1, a.Value())
	assert.Equal(t, (*Item)(nil), a.Next())
	assert.Equal(t, a, b)
	assert.Equal(t, 1, list.Len())
}

func TestPushFrontToNotEmptyList(t *testing.T) {
	list := List{}

	list.PushFront(1)
	b := list.First()

	list.PushFront(2)
	a := list.First()

	assert.Equal(t, 2, a.Value())
	assert.Equal(t, a, b.Prev())
	assert.Equal(t, (*Item)(nil), b.Next())
	assert.Equal(t, b, a.Next())
	assert.Equal(t, (*Item)(nil), a.Prev())
	assert.NotEqual(t, &a, &b)
	assert.Equal(t, 2, list.Len())
}

func TestPushBackToEmptyList(t *testing.T) {
	list := List{}

	list.PushBack(1)

	a := list.Last()
	b := list.First()

	assert.Equal(t, 1, a.Value())
	assert.Equal(t, (*Item)(nil), a.Prev())
	assert.Equal(t, a, b)
	assert.Equal(t, 1, list.Len())
}

func TestPushBackToNotEmptyList(t *testing.T) {
	list := List{}

	list.PushBack(1)
	a := list.First()

	list.PushBack(2)
	b := list.Last()

	assert.Equal(t, 2, b.Value())
	assert.Equal(t, a, b.Prev())
	assert.Equal(t, (*Item)(nil), b.Next())
	assert.Equal(t, b, a.Next())
	assert.Equal(t, (*Item)(nil), a.Prev())
	assert.Equal(t, 2, list.Len())
}

func TestRemoveFromMiddle(t *testing.T) {
	list := List{}

	list.PushBack(1)
	a := list.First()

	list.PushBack(2)
	b := list.Last()

	list.PushBack(3)
	c := list.Last()

	list.Remove(b)

	assert.Equal(t, a, c.Prev())
	assert.Equal(t, c, a.Next())
	assert.Equal(t, 2, list.Len())
}

func TestRemoveFromStart(t *testing.T) {
	list := List{}

	list.PushBack(1)
	a := list.First()

	list.PushBack(2)
	b := list.Last()

	list.PushBack(3)

	list.Remove(a)

	assert.Equal(t, (*Item)(nil), b.Prev())
	assert.Equal(t, b, list.First())
	assert.Equal(t, 2, list.Len())
}

func TestRemoveFromEnd(t *testing.T) {
	list := List{}

	list.PushBack(1)

	list.PushBack(2)
	b := list.Last()

	list.PushBack(3)
	c := list.Last()

	list.Remove(c)

	assert.Equal(t, (*Item)(nil), b.Next())
	assert.Equal(t, b, list.Last())
	assert.Equal(t, 2, list.Len())
}

func TestRemoveNotFromList(t *testing.T) {
	list := List{}

	list.PushBack(1)
	a := list.First()

	list.PushBack(2)
	b := list.Last()

	list.PushBack(3)
	c := list.Last()

	d := Item{
		value: 2,
		next:  c,
		prev:  a,
	}

	list.Remove(&d)

	assert.Equal(t, b, a.Next())
	assert.Equal(t, b, c.Prev())
	assert.Equal(t, 3, list.Len())
}
