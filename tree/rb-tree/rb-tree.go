package rb_tree

import (
	"fmt"
)

type RBTree struct {
	Root     *RBNode
	NodeNum  int64
}



func (t *RBTree)Find(value RBValue) (targetNode *RBNode) {
	targetNode,_ = t.Root.find(value)
	return targetNode
}

func (t *RBTree)newNode(value RBValue) *RBNode {
	node := RBNode{
		Parent: nil,
		LeftChild: nil,
		RightChild: nil,
		nodeColor: red,
		Tree: t,
		Value: value,
	}
	return &node
}

func (t *RBTree)Insert(value RBValue) error {
	node := t.newNode(value)
	if t.Root == nil {
		node.convertColor()
		t.Root = node
		t.NodeNum = 1
		return nil
	}
	targetNode, parentNode := t.Root.find(value)
	if targetNode != nil {
		targetNode.Value.DeepCopy(value)
		return nil
	}
	err := parentNode.insertChild(node)
	if err != nil {
		return err
	}
	t.NodeNum += 1
	return nil
}

func (t *RBTree)Delete(value RBValue) error {
	valueStr, valueErr := value.Marshal()
	if valueErr != nil {
		err := NewRBTreeError(t.Root)
		msg := fmt.Sprintf("delete value failed, invalid RBValue, Marshal failed:%s", valueErr.Error())
		return err.WithMsg(msg)
	}
	if t.Root == nil {
		msg := fmt.Sprintf("delete value:%s in empty tree", string(valueStr))
		return NewEmptyTreeError(msg)
	}
	targetNode, _ := t.Root.find(value)
	if targetNode == nil {
		msg := fmt.Sprintf("delete value:%s can not find in tree", string(valueStr))
		err := NewRBTreeError(t.Root)
		return err.WithMsg(msg)
	}
	var parentNode *RBNode
	if targetNode.LeftChild == nil && targetNode.RightChild == nil {
		return targetNode.removeSelf()
	} else if targetNode.LeftChild != nil {

		_, parentNode = targetNode.LeftChild.find(value)
	} else {
		_, parentNode = targetNode.RightChild.find(value)
	}
	targetNode.Value.DeepCopy(parentNode.Value)
	err := parentNode.removeSelf()
	if err != nil {
		return err
	}
	return nil
}

func (t *RBTree)convertArray() [][]string {
	if t.Root == nil {
		return nil
	}
	treeHeight := t.Root.getTreeHeight()
	totalHeight := treeHeight*2 - 1
	var totalWidth int
	if treeHeight == 1 {
		totalWidth = 1

	} else {
		totalWidth = (2<<(treeHeight-2))*3
	}
	var array = make([][]string, totalHeight)
	for i := range array {
		array[i] = make([]string, totalWidth)
		for j := range array[i] {
			array[i][j] = "   "
		}
	}
	t.Root.convertArray(0, totalWidth / 2, treeHeight,totalWidth / 4, array)
	return array
}


func (t *RBTree)PrintTree() {
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
		Root: nil,
		NodeNum: 0,
	}
	return &tree
}