package captcha

import (
	"reflect"
	"testing"
)

const placeholder = 0

func Test_rightOperand_whenPatternIs1AndRightOperandIs1_rightOperandShouldReturnOne(t *testing.T) {
	c := New(1, placeholder, placeholder, 1)

	l := c.RightOperand()

	if l != "One" {
		t.Errorf("Should be %s but got %s", "One", l)
	}
}

func Test_rightOperand_whenPatternIs1AndRightOperandIs2_rightOperandShouldReturnTwo(t *testing.T) {
	c := New(1, placeholder, placeholder, 2)

	l := c.RightOperand()

	if l != "Two" {
		t.Errorf("Should be %s but got %s", "Two", l)
	}
}

func Test_rightOperand_whenPatternIs1AndRightOperandIs3_rightOperandShouldReturnThree(t *testing.T) {
	c := New(1, placeholder, placeholder, 3)

	r := c.RightOperand()

	if r != "Three" {
		t.Errorf("Should be %s but got %s", "Three", r)
	}
}

func Test_rightOperand_whenPatternIs2AndRightOperandIs1_rightOperandShouldReturn1(t *testing.T) {
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

func Test_operator_shouldCallStringOnce(t *testing.T) {
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

func Test_New_operatorShouldBeOperator(t *testing.T) {
	c := New(placeholder, placeholder, placeholder, placeholder)

	operatorType := reflect.TypeOf(c.operator)
	typeName := operatorType.Name()
	if typeName != "Operator" {
		t.Errorf("c.operator should be Operator but got %s", typeName)
	}
}

func Test_New_whenPatternIs2_leftOperandShouldBeWordOperand(t *testing.T) {
	pattern := 2
	c := New(pattern, placeholder, placeholder, placeholder)

	wordOperandType := reflect.TypeOf(c.leftOperand)
	typeName := wordOperandType.Name()
	if typeName != "WordOperand" {
		t.Errorf("c.leftOperand should be WordOperand but got %s", typeName)
	}
}

func Test_New_whenPatternIs1_rightOperandShouldBeWordOperand(t *testing.T) {
	pattern := 1
	c := New(pattern, placeholder, placeholder, placeholder)

	rightOperandType := reflect.TypeOf(c.rightOperand)
	rightOperandTypeName := rightOperandType.Name()

	if rightOperandTypeName != "WordOperand" {
		t.Errorf("c.rightOperand should be WordOperand but got %s", rightOperandTypeName)
	}
}

func Test_leftOperand_shouldCallStringOnce(t *testing.T) {
	spy := StringSpy{}
	c := New(placeholder, placeholder, placeholder, placeholder)
	c.leftOperand = &spy

	_ = c.LeftOperand()

	if spy.stringCalledCount != 1 {
		t.Errorf("c.leftOperand should call String() once, but call %d", spy.stringCalledCount)
	}
}

func Test_rightOperand_shouldCallStringOnce(t *testing.T) {
	spy := StringSpy{}
	sut := New(placeholder, placeholder, placeholder, placeholder)
	sut.rightOperand = &spy

	sut.RightOperand()

	if spy.stringCalledCount != 1 {
		t.Errorf("sut.rightOperand should call String() once, but call %d", spy.stringCalledCount)
	}
}
