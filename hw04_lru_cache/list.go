package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	CacheKey Key
	Value    interface{}
	Next     *ListItem
	Prev     *ListItem
}

type list struct {
	len   int
	front *ListItem
	back  *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.front
}

func (l *list) Back() *ListItem {
	return l.back
}

func (l *list) PushFront(v interface{}) *ListItem {
	if l.front != nil {
		oldFrontItem := l.front
		l.front = &ListItem{
			Value: v,
			Next:  oldFrontItem,
		}

		oldFrontItem.Prev = l.front

		if l.len == 1 {
			l.back = oldFrontItem
		}

		l.len++
		return l.front
	}

	l.front = &ListItem{Value: v}
	l.len++

	return l.front
}

func (l *list) PushBack(v interface{}) *ListItem {
	if l.back != nil {
		oldBackItem := l.back
		l.back = &ListItem{
			Value: v,
			Prev:  oldBackItem,
		}

		oldBackItem.Next = l.back

		l.len++
		return l.back
	}

	newListItem := &ListItem{Value: v}

	if l.front != nil {
		l.front.Next = newListItem
		l.back = newListItem
		l.back.Prev = l.front
	} else {
		l.front = newListItem
	}

	l.len++

	return newListItem
}

func (l *list) Remove(i *ListItem) {
	i.Prev.Next = i.Next
	if i.Next != nil {
		i.Next.Prev = i.Prev
	} else {
		l.back = i.Prev
	}

	l.len--
}

func (l *list) MoveToFront(i *ListItem) {
	if i == l.front {
		return
	}

	if i == l.back {
		l.back.Prev.Next = nil
		l.front.Prev = l.back
		l.back.Next = l.front
		l.front = l.back
		l.back = l.back.Prev

		return
	}

	i.Next.Prev = i.Prev
	l.front.Prev = i
	i.Next = l.front
	i.Prev = nil

	l.front = i
}

func NewList() List {
	return new(list)
}
