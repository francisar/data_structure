package rb_tree

type RBValue interface {
	LessThan(v RBValue) bool
	Equal(v RBValue) bool
	MoreThan(v RBValue) bool
	DeepCopy(v RBValue)
	Marshal() ([]byte, error)
	UnMarshal(str string) error
}
