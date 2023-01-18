package captcha

// NewLeftOperand is new left operand follow by pattern condition
func NewLeftOperand(pattern, rawValue int) Stringable {
	if pattern == 2 {
		return WordOperand(rawValue)
	}
	return IntOperand(rawValue)
}
