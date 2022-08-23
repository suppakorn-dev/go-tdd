package captcha

import "strconv"

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

var numberToString = map[int]string{
	1: "One",
	2: "Two",
	3: "Three",
	4: "Four",
	5: "Five",
	6: "Six",
	7: "Seven",
	8: "Eight",
	9: "Nine",
}

func (c Captcha) LeftOperand() string {
	if c.pattern == 2 {
		if val, ok := numberToString[c.leftOperand]; ok {
			return val
		}
		return ""

	}
	return strconv.Itoa(c.leftOperand)
}
