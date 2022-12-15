package rb_tree

type color bool

const (
	red   color = true
	black color = false
)

func (c color) String() string {
	switch c {
	case red:
		return "R"
	case black:
		return "B"
	default:

	}
	return "unknown"
}

func (c color) ConvertColor() color {
	return !c
}
