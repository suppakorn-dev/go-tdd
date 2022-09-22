package captcha

import "testing"

func Test_wordOperand_whenPatternIs2_andWordOperandIs1_wordOperandShouldReturnOne(t *testing.T) {
	o := WordOperand(1)

	if o.String() != "One" {
		t.Errorf("Should be %s but got %s", "One", o.String())
	}
}

func Test_wordOperand_whenPatternIs2_andWordOperandIs2_wordOperandShouldReturnTwo(t *testing.T) {
	o := WordOperand(2)

	if o.String() != "Two" {
		t.Errorf("Should be %s but got %s", "Two", o.String())
	}
}

func Test_wordOperand_whenPatternIs2_andWordOperandIs3_wordOperandShouldReturnThree(t *testing.T) {
	o := WordOperand(3)

	if o.String() != "Three" {
		t.Errorf("Should be %s but got %s", "Three", o.String())
	}
}
