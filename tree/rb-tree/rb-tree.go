package rb_tree

import (
	"fmt"
	"github.com/francisar/data_structure"
	"sync"
)

type RBTree struct {
	Root    *RBNode
	mutex      sync.RWMutex
	NodeNum int64
}

func (t *RBTree) Find(item  data_structure.OPItem) (targetNode *RBNode) {
	t.mutex.RLock()
	defer t.mutex.RUnlock()
	targetNode, _ = t.Root.find(item)
	return targetNode
}

func (t *RBTree) newNode(item data_structure.OPItem) *RBNode {
	node := RBNode{
		Parent:     nil,
		LeftChild:  nil,
		RightChild: nil,
		nodeColor:  red,
		Tree:       t,
		Item:       item,
	}
	return &node
}

// Insert an Item into RBTree
func (t *RBTree) Insert(item data_structure.OPItem) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	node := t.newNode(item)
	// empty tree, just insert new node to Root
	if t.Root == nil {
		node.convertColor()
		t.Root = node
		t.NodeNum = 1
		return nil
	}
	targetNode, parentNode := t.Root.find(item)
	if targetNode != nil {
		// item already exists, just update it
		targetNode.Item.DeepCopy(item)
		return nil
	}
	// item is not existed in the tree, insert new item to the position  where it should be in the tree
	err := parentNode.insertChild(node)
	if err != nil {
		return err
	}
	t.NodeNum += 1
	return nil
}

// Delete remove the item from RBTree
func (t *RBTree) Delete(item data_structure.OPItem) error {
	t.mutex.Lock()
	defer t.mutex.Unlock()
	itemStr := item.String()
	if t.Root == nil {
		msg := fmt.Sprintf("delete value:%s in empty tree", itemStr)
		return NewEmptyTreeError(msg)
	}
	targetNode, _ := t.Root.find(item)
	if targetNode == nil {
		msg := fmt.Sprintf("delete value:%s can not find in tree", itemStr)
		err := NewRBTreeError(t.Root)
		return err.WithMsg(msg)
	}
	var parentNode *RBNode
	// if item is a leaf node, just remove it
	// if item is not a leaf node, find the max node of the left child tree or min node of the right tree to delete,
	// copy value of  the real deleting node to the aim deleting node
	if targetNode.LeftChild == nil && targetNode.RightChild == nil {
		return targetNode.removeSelf()
	} else if targetNode.LeftChild != nil {
		// the item is not a leaf node, find the max node in the left child tree while it's not nil
		_, parentNode = targetNode.LeftChild.find(item)
	} else {
		// the item is not a leaf node, find the min node in the right child tree while it's not nil
		_, parentNode = targetNode.RightChild.find(item)
	}

	targetNode.Item.DeepCopy(parentNode.Item)
	err := parentNode.removeSelf()
	if err != nil {
		return err
	}
	return nil
}

func (t *RBTree) convertArray() [][]string {
	if t.Root == nil {
		return nil
	}
	treeHeight := t.Root.getTreeHeight()
	totalHeight := treeHeight*2 - 1
	var totalWidth int
	if treeHeight == 1 {
		totalWidth = 1

	} else {
		totalWidth = (2 << (treeHeight - 2)) * 3
	}
	var array = make([][]string, totalHeight)
	for i := range array {
		array[i] = make([]string, totalWidth)
		for j := range array[i] {
			array[i][j] = "   "
		}
	}
	t.Root.convertArray(0, totalWidth/2, treeHeight, totalWidth/4, array)
	return array
}

func (t *RBTree) PrintTree() {
	// 创建数组
	printArray := t.convertArray()
	// 打印
	if printArray != nil {
		for i := range printArray {
			var res string
			for j := range printArray[i] {
				res = res + printArray[i][j]
			}
			fmt.Println(res)
		}
	} else {
		fmt.Println("this is an empty tree")
	}
}

func NewRBTree() *RBTree {
	tree := RBTree{
		Root:    nil,
		NodeNum: 0,
	}
	return &tree
}
