package linked_list

type Item struct {
	next  *Item
	prev  *Item
	value interface{}
}

func (item *Item) Value() interface{} {
	return item.value
}

func (item *Item) Next() *Item {
	return item.next
}

func (item *Item) Prev() *Item {
	return item.prev
}
