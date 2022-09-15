package captcha

import "strconv"

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     Stringable
	rightOperand int
}

func New(pattern, leftOperand, operator, rightOperand int) Captcha {
	return Captcha{
		pattern:      pattern,
		leftOperand:  leftOperand,
		operator:     Operator(operator),
		rightOperand: rightOperand,
	}
}

func (c Captcha) LeftOperand() string {
	if c.pattern == 2 {
		numberToString := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		return numberToString[c.leftOperand-1]
	}

	return strconv.Itoa(c.leftOperand)
}

func (c Captcha) RightOperand() string {
	if c.pattern == 2 {
		return strconv.Itoa(c.rightOperand)
	}

	numberToString := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}

	return numberToString[c.rightOperand-1]
}

func (c Captcha) Operator() string {
	return c.operator.String()
}
