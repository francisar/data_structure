package rb_tree

import "fmt"

type rbTreeError struct {
	node *RBNode
	msg string
}

func (r *rbTreeError)Error() string {
	return fmt.Sprintf("%v %s", r.node, r.msg)
}

func NewRBTreeError(node *RBNode) *rbTreeError {
	return &rbTreeError{
		node: node,
	}
}

func (r *rbTreeError) WithMsg(msg string) *rbTreeError {
	r.msg = msg
	return r
}

type emptyTreeError struct {
	msg string
}

func (r *emptyTreeError)Error() string {
	return fmt.Sprintf("empty tree err:%s", r.msg)
}

func NewEmptyTreeError(msg string) *emptyTreeError {
	return &emptyTreeError{
		msg: msg,
	}
}