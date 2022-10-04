package captcha

func NewLeftOperand(pattern, rawValue int) Stringable {
	return IntOperand(rawValue)
}
