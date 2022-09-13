package captcha

type Operator struct {
	rawValue int
}

func (o Operator) String() string {
	return "+"
}
