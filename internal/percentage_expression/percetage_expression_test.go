package percentage_expression

import "testing"

func TestEvaluatePercetageExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		hasError bool
	}{
		{"10% of 100", 10.00, false},
		{"50% off 100", 50.00, false},
		{"25.5% of 100", 25.5, false},
		{"25.5% off 100", 74.5, false},
	}

	for _, test := range tests {
		res, err := Evaluate(test.input)

		if test.hasError {
			t.Errorf("for input %q, expected an error but got nil", test.input)
		} else {
			if err != nil {
				t.Errorf("for input %q, expected no error but got: %v", test.input, err)
			} else if res != test.expected {
				t.Errorf("for input %q, expected %v, got %v", test.input, test.expected, res)
			}
		}
	}
}
