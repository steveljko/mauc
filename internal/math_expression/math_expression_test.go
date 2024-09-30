package math_expression

import (
	"testing"

	"mauc/internal/tokenizer"
)

func TestToPostfix(t *testing.T) {
	tests := []struct {
		input    []tokenizer.Token
		expected string
		hasError bool
	}{
		{
			input: []tokenizer.Token{
				{
					Type:  tokenizer.NUMBER,
					Value: "3",
				}, {
					Type:  tokenizer.PLUS,
					Value: "+",
				}, {
					Type:  tokenizer.NUMBER,
					Value: "4",
				}, {
					Type:  tokenizer.EOF,
					Value: "",
				},
			},
			expected: "3 4 +",
			hasError: false,
		},
		{
			input: []tokenizer.Token{{
				Type:  tokenizer.NUMBER,
				Value: "3",
			}, {
				Type:  tokenizer.MULTIPLY,
				Value: "*",
			}, {
				Type:  tokenizer.LPAREN,
				Value: "(",
			}, {
				Type:  tokenizer.NUMBER,
				Value: "4",
			}, {
				Type:  tokenizer.PLUS,
				Value: "+",
			}, {
				Type:  tokenizer.NUMBER,
				Value: "5",
			}, {
				Type:  tokenizer.RPAREN,
				Value: ")",
			}, {
				Type:  tokenizer.EOF,
				Value: "",
			}},
			expected: "3 4 5 + *",
			hasError: false,
		},
		{
			input: []tokenizer.Token{{
				Type:  tokenizer.LPAREN,
				Value: "(",
			}, {
				Type:  tokenizer.NUMBER,
				Value: "3",
			}, {
				Type:  tokenizer.PLUS,
				Value: "+",
			}, {
				Type:  tokenizer.NUMBER,
				Value: "4",
			}, {
				Type:  tokenizer.RPAREN,
				Value: ")",
			}, {
				Type:  tokenizer.EOF,
				Value: "",
			}},
			expected: "3 4 +",
			hasError: false,
		},
		{
			input: []tokenizer.Token{{
				Type:  tokenizer.NUMBER,
				Value: "3",
			}, {
				Type:  tokenizer.RPAREN,
				Value: ")",
			}, {
				Type:  tokenizer.EOF,
				Value: "",
			}},
			expected: "",
			hasError: true,
		},
	}

	for _, test := range tests {
		result, err := ToPostfix(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("for input %v, expected an error but got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for input %q: %v", test.input, err)
			} else if result != test.expected {
				t.Errorf("for input %v, expected %v, got %v", test.input, result, test.expected)
			}
		}
	}
}

func TestEvaluatePostfix(t *testing.T) {
	tests := []struct {
		postfix  string
		expected float64
		hasError bool
	}{
		{"3 4 +", 7, false},
		{"10 2 -", 8, false},
		{"2 3 *", 6, false},
		{"8 2 /", 4, false},
		{"2 3 ^", 8, false},
		{"10 3 %", 1, false},
		{"3 4 + 2 *", 14, false},
		{"5 1 2 + 4 * + 3 -", 14, false},
		{"2 3 4 * +", 14, false},
		{"2 0 /", 0, true},
		{"3 4 + 5", 0, true},
		{"3 +", 0, true},
		{"3 4 5 +", 0, true},
		{"3 4 + 5 * 2 - 1", 0, true},
		{"3 4 + 5 * 2 - 1 2 +", 0, true},
		{"3 4 + 5 * 2 - 1 2 + 3 /", 0, true},
	}

	for _, test := range tests {
		res, err := EvaluatePostfix(test.postfix)
		if test.hasError {
			if err == nil {
				t.Errorf("expected error for postfix %q, got nil", test.postfix)
			}
		} else {
			if err != nil {
				t.Errorf("unexpected error for postfix %q: %v", test.postfix, err)
			} else if res != test.expected {
				t.Errorf("for postfix %q, expected %v, got %v", test.postfix, test.expected, res)
			}
		}
	}
}

func TestEvaluateMathExpression(t *testing.T) {
	tests := []struct {
		input    string
		expected float64
		hasError bool
	}{
		{"3+2", 5, false},
		{"3+(3-2)", 4, false},
		{"10 * 2 + 3", 23, false},
		{"(3*3)^2", 81, false},
		{"(-3*3)*2", -18, false},
		{"(2.5 * 2) + 10", 15, false},
		{"10 % 3", 1, false},
		{"5 / 0", 0, true},
		{"(1 + 2", 0, true},
		{"1 + ) 2", 0, true},
		{"1 + a", 0, true},
	}

	for _, test := range tests {
		res, err := Evaluate(test.input)
		if test.hasError {
			if err == nil {
				t.Errorf("for input %q, expected an error but got nil", test.input)
			}
		} else {
			if err != nil {
				t.Errorf("for input %q, expected no error but got: %v", test.input, err)
			} else if res != test.expected {
				t.Errorf("for input %q, expected %v, got %v", test.input, test.expected, res)
			}
		}
	}
}
