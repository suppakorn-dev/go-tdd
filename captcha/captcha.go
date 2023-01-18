package captcha

// Captcha is struct of pattern, operator and operand
type Captcha struct {
	pattern      int
	leftOperand  Stringable
	operator     Stringable
	rightOperand Stringable
}

// New is create a new captcha object
func New(pattern, leftOperand, operator, rightOperand int) Captcha {
	return Captcha{
		pattern:      pattern,
		leftOperand:  NewLeftOperand(pattern, leftOperand),
		operator:     Operator(operator),
		rightOperand: NewRightOperand(pattern, rightOperand),
	}
}

// LeftOperand will handle left operand
func (c Captcha) LeftOperand() string {
	return c.leftOperand.String()
}

// RightOperand will handle right operand
func (c Captcha) RightOperand() string {
	return c.rightOperand.String()
}

// Operator will handle right operand
func (c Captcha) Operator() string {
	return c.operator.String()
}
