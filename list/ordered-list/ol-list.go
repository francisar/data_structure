package odered_list

import "github.com/francisar/data_structure"

type OrderedList struct {
	Header *OrderedListNode
	Count uint64
}

func (l *OrderedList) newNode(item data_structure.OPItem) *OrderedListNode {
	node := OrderedListNode{
		Next: nil,
		Item: item,
		Pre: nil,
	}
	return &node
}

func NewOrderedList() *OrderedList {
	return &OrderedList{
		Header: nil,
		Count:  0,
	}
}

func (l *OrderedList)Insert(item data_structure.OPItem) {
	node := l.newNode(item)
	if l.Header == nil {
		l.Header = node
		l.Count ++
	}
	pre,current,next := l.Header.find(node)
	if current != nil {
		current.Item.DeepCopy(item)
	}
	l.Count ++
	if pre != nil {
		pre.Next = node
		node.Pre = pre
	}
	if next != nil {
		next.Pre = node
		node.Next = node
	}
	if next == l.Header {
		l.Header = node
	}
}

func (l *OrderedList)Find(item data_structure.OPItem) *OrderedListNode {
	if l.Header == nil {
		return nil
	}
	node := l.newNode(item)
	_,current,_ := l.Header.find(node)
	return current
}

func (l *OrderedList)Delete(item data_structure.OPItem) error {
	if l.Header == nil {
		return nil
	}
	node := l.newNode(item)
	pre,current,next := l.Header.find(node)
	if current != nil {
		if next!=nil {
			next.Pre = pre
		}
		if pre != nil {
			pre.Next = next
		}
		if l.Header == current {
			l.Header = next
		}
	}
}
