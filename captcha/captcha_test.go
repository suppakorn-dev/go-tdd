package captcha

import "testing"

func Test_leftOperand_whenPatternIs1_andLeftOperandIs1_leftOperandShouldReturn1(t *testing.T) {
	placeholder := 0
	c := Captcha{
		pattern:      1,
		leftOperand:  1,
		operator:     placeholder,
		rightOperand: placeholder,
	}

	l := c.LeftOperand()

	if l != "1" {
		t.Error("Should be 1 but got", l)
	}
}

func Test_leftOperand_whenPatternIs1_andLeftOperandIs2_leftOperandShouldReturn2(t *testing.T) {
	placeholder := 0
	c := Captcha{
		pattern:      1,
		leftOperand:  2,
		operator:     placeholder,
		rightOperand: placeholder,
	}

	l := c.LeftOperand()

	if l != "2" {
		t.Error("Should be 2 but got", l)
	}
}

func Test_leftOperand_whenPatternIs2_andLeftOperandIs1_leftOperandShouldReturnOne(t *testing.T) {
	placeholder := 0
	c := Captcha{
		pattern:      2,
		leftOperand:  1,
		operator:     placeholder,
		rightOperand: placeholder,
	}

	l := c.LeftOperand()

	if l != "One" {
		t.Errorf("Should be %s but got %s", "One", l)
	}
}

func Test_leftOperand_whenPatternIs2_andLeftOperandIs2_leftOperandShouldReturnTwo(t *testing.T) {
	placeholder := 0
	c := Captcha{
		pattern:      2,
		leftOperand:  2,
		operator:     placeholder,
		rightOperand: placeholder,
	}

	l := c.LeftOperand()

	if l != "Two" {
		t.Errorf("Should be %s but got %s", "Two", l)
	}
}

func Test_leftOperand_whenPatternIs2_andLeftOpeandIs3_ledtOperandShouldReturnThree(t *testing.T) {
	placeholder := 0
	c := Captcha{
		pattern:      2,
		leftOperand:  3,
		operator:     placeholder,
		rightOperand: placeholder,
	}

	l := c.LeftOperand()

	if l != "Three" {
		t.Errorf("Should be %s but got %s", "Three", l)
	}
}
