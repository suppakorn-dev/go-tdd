package captcha

// WordOperand is alias of int
type WordOperand int

// String is a method to convert type int to string
func (o WordOperand) String() string {
	numberToString := []string{"One", "Two", "Three", "Four", "Five", "Six", "Seven", "Eight", "Nine"}
	return numberToString[o-1]
}
