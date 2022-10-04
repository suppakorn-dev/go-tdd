package captcha

import "strconv"

type IntOperand int

func (o IntOperand) String() string {
	return strconv.Itoa(int(o))
}
