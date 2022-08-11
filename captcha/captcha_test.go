package captcha

import "testing"

func Test_leftOperand_whenPatternIs1_andLeftOperandIs1_shouldReturn1(t *testing.T) {
	c := Captcha{
		pattern:      1,
		leftOperand:  1,
		operator:     0,
		rightOperand: 0,
	}

	l := c.LeftOperand()

	if l != "1" {
		t.Error("Should be 1 but got", l)
	}
}

func Test_leftOperand_whenPatternIs1_andLeftOperandIs2_ShouldReturn2(t *testing.T) {
	c := Captcha{
		pattern:      1,
		leftOperand:  2,
		operator:     0,
		rightOperand: 0,
	}

	l := c.LeftOperand()

	if l != "2" {
		t.Error("Should be 2 but got", l)
	}
}

func Test_leftOperand_whenPatternIs2_andLeftOerandIs1_leftOperandShouldReturnOne(t *testing.T) {
	c := Captcha{
		pattern:      2,
		leftOperand:  1,
		operator:     0,
		rightOperand: 0,
	}

	l := c.LeftOperand()

	if l != "One" {
		t.Errorf("Should be %s but got %s", "One", l)
	}
}
