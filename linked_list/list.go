package linked_list

type List struct {
	first *Item
	last  *Item
	len   int
}

func (list *List) Len() int {
	return list.len
}

func (list *List) First() *Item {
	return list.first
}

func (list *List) Last() *Item {
	return list.last
}

func (list *List) PushFront(value interface{}) {
	first := list.first

	item := Item{
		value: value,
		next:  first,
	}

	list.first = &item

	if first != nil {
		first.prev = &item
	}

	if list.last == nil {
		list.last = &item
	}

	list.len++
}

func (list *List) PushBack(value interface{}) {
	last := list.last

	item := Item{
		value: value,
		prev:  last,
	}

	list.last = &item

	if last != nil {
		last.next = &item
	}

	if list.first == nil {
		list.first = &item
	}

	list.len++
}

func (list *List) Remove(item *Item) {
	i := list.first
	for i != nil && i != item {
		i = i.next
	}

	if i == nil {
		return
	}

	prev := item.Prev()
	next := item.Next()

	if prev != nil {
		prev.next = next
	}

	if next != nil {
		next.prev = prev
	}

	if list.first == item {
		list.first = next
	} else if list.last == item {
		list.last = prev
	}

	list.len--
}
