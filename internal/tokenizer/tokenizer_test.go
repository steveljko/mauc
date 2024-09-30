package tokenizer

import "testing"

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []Token
	}{
		{
			input: "3 + 5",
			expected: []Token{
				{Type: NUMBER, Value: "3"},
				{Type: PLUS, Value: "+"},
				{Type: NUMBER, Value: "5"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "-2 * (3 + 4)",
			expected: []Token{
				{Type: NUMBER, Value: "-2"},
				{Type: MULTIPLY, Value: "*"},
				{Type: LPAREN, Value: "("},
				{Type: NUMBER, Value: "3"},
				{Type: PLUS, Value: "+"},
				{Type: NUMBER, Value: "4"},
				{Type: RPAREN, Value: ")"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input: "10.5 / 2",
			expected: []Token{
				{Type: NUMBER, Value: "10.5"},
				{Type: DIVIDE, Value: "/"},
				{Type: NUMBER, Value: "2"},
				{Type: EOF, Value: ""},
			},
		},
		{
			input:    "x + y",
			expected: nil,
		},
	}

	for _, test := range tests {
		tokens, err := Tokenize(test.input)
		if test.expected == nil {
			if err == nil {
				t.Errorf("Expected an error for input %q, but got none", test.input)
			}
			continue
		}

		if err != nil {
			t.Errorf("Unexpected error for input %q: %v", test.input, err)
			continue
		}

		if len(tokens) != len(test.expected) {
			t.Errorf("For input %q, expected %d tokens, but got %d", test.input, len(test.expected), len(tokens))
			continue
		}

		for i, token := range tokens {
			if token != test.expected[i] {
				t.Errorf("For input %q, expected token %v, but got %v", test.input, test.expected[i], token)
			}
		}
	}
}
