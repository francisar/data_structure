package data_structure



type OPItem interface {
	LessThan(v OPItem) bool
	Equal(v OPItem) bool
	MoreThan(v OPItem) bool
	DeepCopy(v OPItem)
	Marshal() ([]byte, error)
	UnMarshal(str string) error
	String() string
}
