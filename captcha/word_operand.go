package captcha

type WordOperand int

func (o WordOperand) String() string {
	numberToString := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	return numberToString[o-1]
}
