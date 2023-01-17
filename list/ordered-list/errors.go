package ordered_list


import "fmt"

type orderedlISTError struct {
	node *OrderedListNode
	msg  string
}

func (r *orderedlISTError) Error() string {
	return fmt.Sprintf("%s %s", r.node.Item.String(), r.msg)
}

func NewOrderedListError(node *OrderedListNode) *orderedlISTError {
	return &orderedlISTError{
		node: node,
	}
}

func (r *orderedlISTError) WithMsg(msg string) *orderedlISTError {
	r.msg = msg
	return r
}

type emptyOrderedListError struct {
	msg string
}

func (r *emptyOrderedListError) Error() string {
	return fmt.Sprintf("empty ordered list err:%s", r.msg)
}

func NewEmptyOrderedListError(msg string) *emptyOrderedListError {
	return &emptyOrderedListError{
		msg: msg,
	}
}