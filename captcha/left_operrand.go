package captcha

func NewLeftOperand(pattern, rawValue int) Stringable {
	if pattern == 2 {
		return WordOperand(rawValue)
	}
	return IntOperand(rawValue)
}
