package captcha

type Captcha struct {
	pattern      int
	leftOperand  Stringable
	operator     Stringable
	rightOperand int
}

func New(pattern, leftOperand, operator, rightOperand int) Captcha {
	return Captcha{
		pattern:      pattern,
		leftOperand:  NewLeftOperand(pattern, leftOperand),
		operator:     Operator(operator),
		rightOperand: rightOperand,
	}
}

func (c Captcha) LeftOperand() string {
	return c.leftOperand.String()
}

func (c Captcha) RightOperand() string {
	if c.pattern == 2 {
		return IntOperand(c.rightOperand).String()
	}

	return WordOperand(c.rightOperand).String()
}

func (c Captcha) Operator() string {
	return c.operator.String()
}
