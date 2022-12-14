package captcha

// NewRightOperand is new right operand follow by pattern condition
func NewRightOperand(pattern int, rawValue int) Stringable {
	if pattern == 1 {
		return WordOperand(rawValue)
	}

	return IntOperand(rawValue)
}
