package captcha

import "testing"

func TestNewLeftOperandWhenPatternIs1ShouldReturnIntOperand(t *testing.T) {
	pattern := 1
	leftOperand := 1

	sut := NewLeftOperand(pattern, leftOperand)

	_, ok := sut.(IntOperand)
	if !ok {
		t.Errorf("Expect type IntOperand but got %T", sut)
	}
}

func Test_NewLeftOperand_whenPatternIs2_shouldReturnWordOperand(t *testing.T) {
	pattern := 2
	leftOperand := 1

	sut := NewLeftOperand(pattern, leftOperand)

	_, ok := sut.(WordOperand)
	if !ok {
		t.Errorf("Expect type WordOperand but got %T", sut)
	}
}
