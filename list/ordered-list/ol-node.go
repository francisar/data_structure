package odered_list

import "github.com/francisar/data_structure"

type OrderedListNode struct {
	Pre *OrderedListNode
	Next *OrderedListNode
	Item data_structure.OPItem
}

func (n *OrderedListNode)find(node *OrderedListNode) (pre *OrderedListNode,current *OrderedListNode, next *OrderedListNode) {
	if n.Item.LessThan(node.Item) {
		if n.Next != nil {
			return n.Next.find(node)
		} else {
			return n,nil, nil
		}
	} else if n.Item.Equal(node.Item) {
		return n.Pre, n, n.Next
	} else {
		return n.Pre, nil, n
	}
}
