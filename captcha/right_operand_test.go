package captcha

import (
	"reflect"
	"testing"
)

func Test_NewRightOperand_whenPatternIs2_shouldReturnIntOperand(t *testing.T) {
	pattern := 2
	rightOperand := 1

	sut := NewRightOperand(pattern, rightOperand)

	rightOperandType := reflect.TypeOf(sut)
	rightOperandTypeName := rightOperandType.Name()

	if rightOperandTypeName != "IntOperand" {
		t.Errorf("Expect type IntOperand but got %s", rightOperandTypeName)
	}
}

func Test_NewRightOperand_whenPatternIs1_shouldReturnWordOperand(t *testing.T) {
	pattern := 1
	rightOperand := 1

	sut := NewRightOperand(pattern, rightOperand)

	rightOperandType := reflect.TypeOf(sut)
	rightOperandTypeName := rightOperandType.Name()

	if rightOperandTypeName != "WordOperand" {
		t.Errorf("Expect type WordOperand but got %s", rightOperandTypeName)
	}
}
