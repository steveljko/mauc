package tokenizer

import (
	"testing"

	"mauc/internal/tokenizer"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		input    string
		expected []tokenizer.Token
	}{
		{
			input: "3 + 5",
			expected: []tokenizer.Token{
				{Type: tokenizer.NUMBER, Value: "3"},
				{Type: tokenizer.PLUS, Value: "+"},
				{Type: tokenizer.NUMBER, Value: "5"},
				{Type: tokenizer.EOF, Value: ""},
			},
		},
		{
			input: "-2 * (3 + 4)",
			expected: []tokenizer.Token{
				{Type: tokenizer.NUMBER, Value: "-2"},
				{Type: tokenizer.MULTIPLY, Value: "*"},
				{Type: tokenizer.LPAREN, Value: "("},
				{Type: tokenizer.NUMBER, Value: "3"},
				{Type: tokenizer.PLUS, Value: "+"},
				{Type: tokenizer.NUMBER, Value: "4"},
				{Type: tokenizer.RPAREN, Value: ")"},
				{Type: tokenizer.EOF, Value: ""},
			},
		},
		{
			input: "10.5 / 2",
			expected: []tokenizer.Token{
				{Type: tokenizer.NUMBER, Value: "10.5"},
				{Type: tokenizer.DIVIDE, Value: "/"},
				{Type: tokenizer.NUMBER, Value: "2"},
				{Type: tokenizer.EOF, Value: ""},
			},
		},
		{
			input:    "x + y",
			expected: nil,
		},
	}

	for _, test := range tests {
		tokens, err := tokenizer.Tokenize(test.input)
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
