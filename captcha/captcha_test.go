package captcha

import "testing"

const placeholder = 0

func Test_leftOperand_whenPatternIs1_andLeftOperandIs1_leftOperandShouldReturn1(t *testing.T) {
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

func Test_rightOperand_whenPatternIs1_andRightOperandIs1_rightOperandShouldReturnOne(t *testing.T) {
	c := Captcha{
		pattern:      1,
		leftOperand:  placeholder,
		operator:     placeholder,
		rightOperand: 1,
	}

	l := c.RightOperand()

	if l != "One" {
		t.Errorf("Should be %s but got %s", "One", l)
	}
}

func Test_rightOperand_whenPatternIs1_andRightOperandIs2_rightOperandShouldReturnTwo(t *testing.T) {
	c := Captcha{
		pattern:      1,
		leftOperand:  placeholder,
		operator:     placeholder,
		rightOperand: 2,
	}

	l := c.RightOperand()

	if l != "Two" {
		t.Errorf("Should be %s but got %s", "Two", l)
	}
}

func Test_rightOperand_whenPatternIs1_andRightOperandIs3_rightOperandShouldReturnThree(t *testing.T) {
	c := Captcha{
		pattern:      1,
		leftOperand:  placeholder,
		operator:     placeholder,
		rightOperand: 3,
	}

	r := c.RightOperand()

	if r != "Three" {
		t.Errorf("Should be %s but got %s", "Three", r)
	}
}

func Test_rightOperand_whenPatternIs2_andRightOperandIs1_rightOperandShouldReturn1(t *testing.T) {
	c := Captcha{
		pattern:      2,
		leftOperand:  placeholder,
		operator:     placeholder,
		rightOperand: 1,
	}

	r := c.RightOperand()

	if r != "1" {
		t.Errorf("Should be %s but got %s", "1", r)
	}
}
