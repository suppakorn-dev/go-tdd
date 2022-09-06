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

	if c.rightOperand == 3 {
		numberToString := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		return numberToString[c.rightOperand-1]
	}

	if c.rightOperand == 2 {
		numberToString := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
		return numberToString[c.rightOperand-1]
	}

	numberToString := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	return numberToString[c.rightOperand-1]
}
