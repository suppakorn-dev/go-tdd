package captcha

type Captcha struct {
	pattern      int
	leftOperand  int
	operator     int
	rightOperand int
}

func (c Captcha) LeftOperand() int {
	return 1
}
