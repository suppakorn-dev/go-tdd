package captcha

import "testing"

func Test_IntOperandIs1_String_should1(t *testing.T) {
	o := IntOperand(1)

	if o.String() != "1" {
		t.Errorf("Should be 1 but got %s", o.String())
	}
}

func Test_IntOperandIs2_String_should2(t *testing.T) {
	o := IntOperand(2)

	if o.String() != "2" {
		t.Errorf("Should be 2 but got %s", o.String())
	}
}

func Test_IntOperandIs3_String_should3(t *testing.T) {
	o := IntOperand(3)

	if o.String() != "3" {
		t.Errorf("Should be 3 but got %s", o.String())
	}
}
