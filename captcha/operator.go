package captcha

// Operator is type to operate string
// Operator
type Operator int

func (o Operator) String() string {
	if o == 3 {
		return "/"
	}

	if o == 2 {
		return "-"
	}

	return "+"
}
