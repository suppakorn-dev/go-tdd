package captcha

import "strconv"

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
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
	if c.operator == 3 {
		return Operator(c.operator).String()
	}

	if c.operator == 2 {
		return Operator(c.operator).String()
	}

	return Operator(c.operator).String()
}
