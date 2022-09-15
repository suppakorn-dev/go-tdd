package captcha

import (
	"testing"
)

const placeholder = 0

func Test_leftOperand_whenPatternIs1_andLeftOperandIs1_leftOperandShouldReturn1(t *testing.T) {
	c := New(1, 1, placeholder, placeholder)

	l := c.LeftOperand()

	if l != "1" {
		t.Error("Should be 1 but got", l)
	}
}

func Test_leftOperand_whenPatternIs1_andLeftOperandIs2_leftOperandShouldReturn2(t *testing.T) {
	c := New(1, 2, placeholder, placeholder)

	l := c.LeftOperand()

	if l != "2" {
		t.Error("Should be 2 but got", l)
	}
}

func Test_leftOperand_whenPatternIs2_andLeftOperandIs1_leftOperandShouldReturnOne(t *testing.T) {
	c := New(2, 1, placeholder, placeholder)

	l := c.LeftOperand()

	if l != "One" {
		t.Errorf("Should be %s but got %s", "One", l)
	}
}

func Test_leftOperand_whenPatternIs2_andLeftOperandIs2_leftOperandShouldReturnTwo(t *testing.T) {
	c := New(2, 2, placeholder, placeholder)

	l := c.LeftOperand()

	if l != "Two" {
		t.Errorf("Should be %s but got %s", "Two", l)
	}
}

func Test_leftOperand_whenPatternIs2_andLeftOperandIs3_leftOperandShouldReturnThree(t *testing.T) {
	c := New(2, 3, placeholder, placeholder)

	l := c.LeftOperand()

	if l != "Three" {
		t.Errorf("Should be %s but got %s", "Three", l)
	}
}

func Test_rightOperand_whenPatternIs1_andRightOperandIs1_rightOperandShouldReturnOne(t *testing.T) {
	c := New(1, placeholder, placeholder, 1)

	l := c.RightOperand()

	if l != "One" {
		t.Errorf("Should be %s but got %s", "One", l)
	}
}

func Test_rightOperand_whenPatternIs1_andRightOperandIs2_rightOperandShouldReturnTwo(t *testing.T) {
	c := New(1, placeholder, placeholder, 2)

	l := c.RightOperand()

	if l != "Two" {
		t.Errorf("Should be %s but got %s", "Two", l)
	}
}

func Test_rightOperand_whenPatternIs1_andRightOperandIs3_rightOperandShouldReturnThree(t *testing.T) {
	c := New(1, placeholder, placeholder, 3)

	r := c.RightOperand()

	if r != "Three" {
		t.Errorf("Should be %s but got %s", "Three", r)
	}
}

func Test_rightOperand_whenPatternIs2_andRightOperandIs1_rightOperandShouldReturn1(t *testing.T) {
	c := New(2, placeholder, placeholder, 1)

	r := c.RightOperand()

	if r != "1" {
		t.Errorf("Should be %s but got %s", "1", r)
	}
}

func Test_operator_whenInputIs1_operatorShouldReturnAdditionalSymbol(t *testing.T) {
	c := New(placeholder, placeholder, 1, placeholder)

	r := c.Operator()

	if r != "+" {
		t.Errorf("Should be %s but got %s", "+", r)
	}
}

func Test_operator_whenInputIs2_operatorShouldReturnMinusSymbol(t *testing.T) {
	c := New(placeholder, placeholder, 2, placeholder)

	o := c.Operator()

	if o != "-" {
		t.Errorf("Should be %s but got %s", "-", o)
	}
}

func Test_operator_whenInputIs3_operatorShouldReturnDivideSymbol(t *testing.T) {
	c := New(placeholder, placeholder, 3, placeholder)

	o := c.Operator()

	if o != "/" {
		t.Errorf("Should be %s but got %s", "/", o)
	}
}

type StringSpy struct {
	mockString        string
	stringCalledCount uint
}

func (s *StringSpy) String() string {
	s.stringCalledCount++
	return s.mockString
}

func Test_operatorShouldCallStringOnce(t *testing.T) {
	c := New(placeholder, placeholder, placeholder, placeholder)
	s := StringSpy{mockString: "placeholder"}
	c.operator = &s

	actual := c.Operator()

	if s.stringCalledCount != 1 {
		t.Errorf("Should call String only once but called %d", s.stringCalledCount)
	}

	if actual != s.mockString {
		t.Errorf("It should return %s but got %s", s.mockString, actual)
	}
}

func Test_operatorShouldBeOperator(t *testing.T) {
	c := New(placeholder, placeholder, placeholder, placeholder)

	_, ok := c.operator.(Operator)

	if !ok {
		t.Errorf("c.operator should be Operator")
	}
}
