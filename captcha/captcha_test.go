package captcha

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

const placeholder = 0

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

func Test_captcha(t *testing.T) {
	captcha := New(1, 1, 1, 1)
	expected := "1 + One"

	actual := result(captcha)

	assert.Equal(t, expected, actual)
}

func result(c Captcha) string {
	return fmt.Sprintf("%v %v %v", c.LeftOperand(), c.Operator(), c.RightOperand())
}
