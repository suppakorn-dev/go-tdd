package captcha

import (
	"reflect"
	"testing"
)

func Test_NewLeftOperand_whenPatternIs1_shouldReturnIntOperand(t *testing.T) {
	pattern := 1
	leftOperand := 1

	sut := NewLeftOperand(pattern, leftOperand)

	leftOperandType := reflect.TypeOf(sut)
	typeName := leftOperandType.Name()
	if typeName != "IntOperand" {
		t.Errorf("Expect type IntOperand but got %s", typeName)
	}
}

func Test_NewLeftOperand_whenPatternIs2_shouldReturnWordOperand(t *testing.T) {
	pattern := 2
	leftOperand := 1

	sut := NewLeftOperand(pattern, leftOperand)

	leftOperandType := reflect.TypeOf(sut)
	typeName := leftOperandType.Name()
	if typeName != "WordOperand" {
		t.Errorf("Expect type WordOperand but got %s", typeName)
	}
}
