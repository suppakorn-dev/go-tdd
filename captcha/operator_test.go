package captcha

import "testing"

func Test_whenOperatorIs1_stringShouldReturnPlus(t *testing.T) {
	var o Operator = 1

	if o.String() != "+" {
		t.Errorf("Operator.String() should return + but got %s", o.String())
	}
}
