package captcha

import "strconv"

// IntOperand is alias of int
type IntOperand int

// String is a method to convert type int to string
func (o IntOperand) String() string {
	return strconv.Itoa(int(o))
}
