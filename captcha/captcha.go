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
		return "One"
	}
	return strconv.Itoa(c.leftOperand)
}
