package rb_tree

import (
	"fmt"
	"github.com/francisar/data_structure/util"
)

type RBNode struct {
	nodeColor color
	LeftChild *RBNode
	RightChild *RBNode
	Parent     *RBNode
	Value RBValue
	Tree *RBTree
}

func (r *RBNode)isBlack() bool {
	return r.nodeColor == black
}

func (r *RBNode)isRed() bool {
	return r.nodeColor == red
}

func (r *RBNode)turnBlack()  {
	r.nodeColor =  black
}

func (r *RBNode)turnRed()  {
	r.nodeColor = red
}

func (r *RBNode)isLeftBlack() bool {
	if r.LeftChild == nil {
		return true
	}
	return r.LeftChild.isBlack()
}

func (r *RBNode)isRightBlack() bool {
	if r.RightChild == nil {
		return true
	}
	return r.RightChild.isBlack()
}

func (r *RBNode)isRoot() bool {
	return r.Parent == nil && r.Tree.Root == r
}

func (r *RBNode)convertColor()  {
	r.nodeColor = r.nodeColor.ConvertColor()
}

func (r *RBNode)isParentLeft() bool {
	if r.isRoot(){
		return false
	}
	return r == r.Parent.LeftChild
}

func (r *RBNode)isParentRight() bool {
	if r.isRoot() {
		return false
	}
	return r == r.Parent.RightChild
}

func (r *RBNode)leftRotation()  {
	right := r.RightChild
	if r.isRoot() {
		r.Tree.Root = right
	} else if r.isParentLeft() {
		r.Parent.LeftChild = right
	} else if r.isParentRight() {
		r.Parent.RightChild = right
	}
	if right != nil {
		right.Parent = r.Parent
		r.RightChild = right.LeftChild
		if r.RightChild != nil {
			r.RightChild.Parent = r
		}
		right.LeftChild = r

	}
	r.Parent = right
}


func (r *RBNode)rightRotation()  {
	left := r.LeftChild
	if r.isRoot() {
		r.Tree.Root = left
	} else if r.isParentLeft() {
		r.Parent.LeftChild = left
	} else if r.isParentRight() {
		r.Parent.RightChild = left
	}
	if left != nil {
		left.Parent = r.Parent
		r.LeftChild = left.RightChild
		if r.LeftChild != nil {
			r.LeftChild.Parent = r
		}
		left.RightChild = r
	}
	r.Parent = left
}

func (r *RBNode)find(value RBValue) (targetNode *RBNode, parentNode *RBNode) {
	if value.Equal(r.Value) {
		return r, r.Parent
	} else if value.LessThan(r.Value) {
		if r.LeftChild == nil {
			return nil, r
		}
		return r.LeftChild.find(value)
	} else {
		if r.RightChild == nil {
			return nil, r
		}
		return r.RightChild.find(value)
	}
}


// insertChild insert child node to current node
// current node must have at least one nil node
func (r *RBNode)insertChild(node *RBNode) error {
	// make sure the node to be inserted is red color
	if node.isBlack() {
		node.convertColor()
	}
	// insert new node as left child of current node if left child is nil node
	if r.Value.LessThan(node.Value) {
		if  r.RightChild != nil {
			err := NewRBTreeError(node)
			return err.WithMsg("insertRightChild while RightChild is not Nil")
		}
		r.RightChild = node
		node.Parent = r
	} else {
		if r.LeftChild != nil {
			err := NewRBTreeError(node)
			return err.WithMsg("insertLeftChild while LeftChild is not Nil")
		}
		r.LeftChild = node
		node.Parent = r
	}
	parentNode := r
	for parentNode.isRed() {
		grandParent := parentNode.Parent
		uncleNode := grandParent.LeftChild
		if parentNode.isParentLeft()  {
			uncleNode = grandParent.RightChild
			if uncleNode == nil {
				if node.isParentRight() {
					r.leftRotation()
					parentNode = node
				}
				parentNode.convertColor()
				grandParent.convertColor()
				grandParent.rightRotation()
				return nil
			}
		} else {
			if uncleNode == nil {
				if node.isParentLeft() {
					r.rightRotation()
					parentNode = node
				}
				parentNode.convertColor()
				grandParent.convertColor()
				grandParent.leftRotation()
				return nil
			}
		}

		parentNode.convertColor()
		if uncleNode.isRed() {
			uncleNode.convertColor()
			if !grandParent.isRoot() {
				grandParent.convertColor()
				parentNode = grandParent.Parent
			}
		} else {
			if uncleNode.isParentRight() {
				grandParent.rightRotation()
			} else if uncleNode.isParentLeft() {
				grandParent.leftRotation()
			}
			grandParent.convertColor()
			parentNode = grandParent.Parent
		}
	}
	return nil
}


//
func (r *RBNode)removeSelf() error {
	if r.LeftChild != nil && r.RightChild != nil {
		err := NewRBTreeError(r)
		return err.WithMsg("try to remove node with two child node ")
	}
	removeNode := r
	//
	if r.LeftChild != nil {
		r.Value.DeepCopy(r.LeftChild.Value)
		removeNode = r.LeftChild
	} else if r.RightChild != nil {
		r.Value.DeepCopy(r.LeftChild.Value)
		removeNode = r.RightChild
	}
	if r.Parent == nil {
		r.Tree.Root = nil
	} else {
		brotherNode := removeNode.Parent.LeftChild
		if removeNode.isParentLeft(){
			removeNode.Parent.LeftChild = nil
			brotherNode = removeNode.Parent.RightChild
		} else {
			removeNode.Parent.RightChild = nil
		}
		if removeNode.isBlack(){
			return brotherNode.removeRebalance()
		}
	}
	r.Tree.NodeNum -= 1
	return nil
}

// removeRebalance is called by removeSelf
// deal with removeNode and brotherNode both are black
func (r *RBNode)removeRebalance() error {
	if r.isRed() {
		// removeNode is black and brotherNode is red,
		//so the parent Node must be black,
		//and brotherNode must have two black children
		r.convertColor()
		r.Parent.convertColor()
		brotherNode := r.LeftChild
		if r.isParentLeft(){
			brotherNode = r.RightChild
			r.Parent.rightRotation()
		} else {
			r.Parent.leftRotation()
		}
		return brotherNode.removeRebalance()
	} else {
		// brotherNode has children, which means the children must be red
		if (r.LeftChild !=nil && r.LeftChild.isRed()) || (r.RightChild != nil && r.RightChild.isRed()) {
			return r.brotherWithRedChild()
		} else {
			return r.doubleBlackBalance()
		}
	}
}

func (r *RBNode)doubleBlackBalance() error {
	// removeNode and brotherNode both are black and brotherNode has no child,
	if r.Parent.isRed() {
		/*
			 case1:
			                    |red(4)|                                                    |black(4)|
			                     /      \                                                     /
			          |black(2)(r)|      |black(5)(need to be removed)|  ->             |red(2)|
			case2:
			                             |red(4)|                                    |black(4)|
			                               /      \                                        \
			   |black(3)(need to be removed)|   |black(5)(r)|           ->                 |red(5)|
			       /
			|red(2)|
		*/
		r.convertColor()
		r.Parent.convertColor()
	} else {
		/*deal with this case
		                     |black(4)|
		                    /      \
		         |black(2)(r)|      |black(5)(need to be removed)|
		or
		                             |black(4)|
		                             /      \
		 |black(3)(need to be removed)|   |black(5)(r)|
		*/
		r.convertColor()
		currentNode := r.Parent
		if currentNode.isRoot() {
			return nil
		}
		brotherNode := currentNode.Parent.LeftChild
		if currentNode.isParentLeft() {
			brotherNode = currentNode.Parent.RightChild
		}
		return brotherNode.removeRebalance()
	}
	return nil
}

func (r* RBNode)brotherWithRedChild() error {
	if r.isParentLeft()  {
		/* brotherNode is left child
		 case1:
		                    |parent(4)|                                                    |parent(2)|
		                     /      \                                                     /           \
		          |black(2)(r)|      |black(5)(need to be removed)|  ->             |black(1)|          |black(4)|
		            /         \                                                                        /
		      |red(1)|        |red(3)|                                                            |red(3)|
		case2:
		           |parent(4)|                                                        |parent(3)|
		            /      \                                                         /            \
		   |black(3)(r)|   |black(5)(need to be removed)|           ->          |black(2)|          |black(4)|
		       /
		|red(2)|
		*/
		rotationNode := r.Parent
		brotherNode := r
		if  r.LeftChild == nil || r.LeftChild.isBlack() {
			/* brotherNode is left child and has a right child, so the acitons in this section as follows:
			            |parent(4)|                                            |parent(4)|
			            /      \                                              /          \
			   |black(2)|     |black(5)(need to be removed)|     ->     |red(3)|         |black(5)(need to be removed)|
			             \                                                  /
			            |red(3)|                                     |black(2)|
			*/
			brotherNode = r.RightChild
			r.convertColor()
			r.RightChild.convertColor()
			r.leftRotation()
		}
		brotherNode.nodeColor = rotationNode.nodeColor
		brotherNode.LeftChild.turnBlack()
		rotationNode.turnBlack()
		rotationNode.rightRotation()
	} else if r.isParentRight(){
		/* brotherNode is right child
		 case1:
		                           |parent(4)|                                            |parent(6)|
		                            /      \                                             /           \
		|black(2)(need to be removed)|      |black(6)(r)|            ->             |black(4)|          |black(7)|
		                                    /         \                                   \
		                              |red(5)|        |red(7)|                            |red(5)|
		case2:
		                            |parent(4)|                                            |parent(6)|
		                            /      \                                             /           \
		|black(2)(need to be removed)|      |black(6)(r)|            ->             |black(4)|          |black(7)|
		                                             \
		                                          |red(7)|
		*/
		rotationNode := r.Parent
		brotherNode := r
		if  r.RightChild == nil || r.RightChild.isBlack(){
			/* brotherNode is right child and has a left child, so the acitons in this section as follows:
			                            |parent(4)|                                             |parent(4)|
			                             /      \                                              /          \
			|black(2)(need to be removed)|     |black(6)(r)|     ->     |black(2)(need to be removed)|       |red(5)|
			                                   /                                                                \
			                               |red(5)|                                                          |black(6)|
			*/
			left := r.LeftChild
			brotherNode = left
			r.convertColor()
			left.convertColor()
			r.rightRotation()
		}
		brotherNode.nodeColor = rotationNode.nodeColor
		brotherNode.RightChild.turnBlack()
		rotationNode.turnBlack()
		rotationNode.leftRotation()
	}
	return nil
}

func (r *RBNode)getTreeHeight() int {
	leftHeight := 0
	rightHeight := 0
	if r.LeftChild != nil {
		leftHeight = r.LeftChild.getTreeHeight()
	}
	if r.RightChild != nil {
		rightHeight = r.RightChild.getTreeHeight()
	}
	return 1 + util.MaxInt(leftHeight, rightHeight)
}


func (r *RBNode)convertArray(row int, column int, treeHeight int, gap int, array [][]string) {
	valueByte,valueErr:= r.Value.Marshal()
	var valueStr string
	if valueErr != nil {
		valueStr = "invalid"
	} else {
		valueStr = string(valueByte)
	}
	if util.IsEven(column) {
		array[row][column] = fmt.Sprintf("%s,%s  ", r.nodeColor.String(), valueStr)
		if column != 0 {
			array[row][column-1] = " "
		}
	} else {
		array[row][column] = fmt.Sprintf("%s,%s", r.nodeColor.String(), valueStr)
	}
	if r.LeftChild != nil {
		slashColumn := column - gap
		if util.IsEven(slashColumn) {
			array[row+1][slashColumn] = " /   "
			if slashColumn !=0 {
				array[row+1][slashColumn-1] = " "
			}
		} else {
			array[row+1][slashColumn] = " / "
		}
		r.LeftChild.convertArray(row+2, slashColumn, treeHeight,gap/2, array)
	}
	if r.RightChild != nil {
		slashColumn := column + gap
		if util.IsEven(slashColumn) {
			array[row+1][slashColumn] = "  \\ "
			if slashColumn == column * 2 {
				array[row+1][slashColumn+1] = " "
			}
		} else {
			array[row+1][slashColumn] = " \\ "
		}
		r.RightChild.convertArray(row+2, slashColumn, treeHeight,gap/2, array)
	}
}