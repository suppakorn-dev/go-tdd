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

	if l != 1 {
		t.Error("Should be 1 but got", l)
	}
}
