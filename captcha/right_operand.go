package captcha

// NewRightOperand is new right operand follow by pattern condition
func NewRightOperand(pattern int, rawValue int) Stringable {
	return IntOperand(rawValue)
}
