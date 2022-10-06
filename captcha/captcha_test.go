package captcha

import (
	"reflect"
	"testing"
)

const placeholder = 0

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

	operatorType := reflect.TypeOf(c.operator)
	typeName := operatorType.Name()
	if typeName != "Operator" {
		t.Errorf("c.operator should be Operator but got %s", typeName)
	}
}

func Test_leftOperandShouldBeWordOperandWhenPatterIs2(t *testing.T) {
	pattern := 2
	c := New(pattern, placeholder, placeholder, placeholder)

	wordOperandType := reflect.TypeOf(c.leftOperand)
	typeName := wordOperandType.Name()
	if typeName != "WordOperand" {
		t.Errorf("c.leftOperand should be WordOperand but got %s", typeName)
	}
}

func Test_leftOperandShouldCallStringOnce(t *testing.T) {
	spy := StringSpy{}
	c := New(placeholder, placeholder, placeholder, placeholder)
	c.leftOperand = &spy

	_ = c.LeftOperand()

	if spy.stringCalledCount != 1 {
		t.Errorf("c.leftOperand should call String() once, but call %d", spy.stringCalledCount)
	}
}
