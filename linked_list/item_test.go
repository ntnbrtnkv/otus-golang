package linked_list

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGetValue(t *testing.T) {
	val := 1
	item := Item{value: val}
	assert.Equal(t, val, item.Value())

	slice := []int{1, 2}
	item = Item{value: slice}
	assert.Equal(t, slice, item.Value())
}

func TestGetNext(t *testing.T) {
	a := Item{value: 1}
	b := Item{value: 2, next: &a}
	assert.Equal(t, a, *b.Next())
}

func TestGetPrev(t *testing.T) {
	a := Item{value: 1}
	b := Item{value: 2, prev: &a}
	assert.Equal(t, a, *b.Prev())
}
