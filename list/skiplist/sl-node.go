package skiplist

import "github.com/francisar/data_structure"

type Node struct {
	Item data_structure.OPItem
	Next []*Node
	Pre *Node
	Level int
}

