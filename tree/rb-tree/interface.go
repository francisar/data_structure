package rb_tree

type RBItem interface {
	LessThan(v RBItem) bool
	Equal(v RBItem) bool
	MoreThan(v RBItem) bool
	DeepCopy(v RBItem)
	Marshal() ([]byte, error)
	UnMarshal(str string) error
	String() string
}
