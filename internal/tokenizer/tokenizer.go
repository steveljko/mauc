package tokenizer

import (
	"fmt"
	"unicode"
)

type TokenType int

const (
	EOF TokenType = iota
	NUMBER
	PLUS
	MINUS
	MULTIPLY
	DIVIDE
	POWER
	MODULO
	LPAREN
	RPAREN
)

type Token struct {
	Type  TokenType
	Value string
}

func Tokenize(input string) ([]Token, error) {
	var tokens []Token
	var currentToken string

	for i := 0; i < len(input); i++ {
		token := input[i]

		// skip whitespace
		if unicode.IsSpace(rune(token)) {
			continue
		}

		switch {
		// check if the character is a digit, a decimal point, or a minus sign indicating a negative number
		case unicode.IsDigit(rune(token)) || (token == '-' && (i == 0 || input[i-1] == '(')):
			currentToken = string(token)
			i++
			for i < len(input) && (unicode.IsDigit(rune(input[i])) || input[i] == '.') {
				currentToken += string(input[i])
				i++
			}
			i--
			tokens = append(tokens, Token{Type: NUMBER, Value: currentToken})
		case token == '+':
			tokens = append(tokens, Token{Type: PLUS, Value: string(token)})
		case token == '-':
			tokens = append(tokens, Token{Type: MINUS, Value: string(token)})
		case token == '*':
			tokens = append(tokens, Token{Type: MULTIPLY, Value: string(token)})
		case token == '/':
			tokens = append(tokens, Token{Type: DIVIDE, Value: string(token)})
		case token == '^':
			tokens = append(tokens, Token{Type: POWER, Value: string(token)})
		case token == '%':
			tokens = append(tokens, Token{Type: MODULO, Value: string(token)})
		case token == '(':
			tokens = append(tokens, Token{Type: LPAREN, Value: string(token)})
		case token == ')':
			tokens = append(tokens, Token{Type: RPAREN, Value: string(token)})
		default:
			return nil, fmt.Errorf("Unexpected character")
		}
	}

	tokens = append(tokens, Token{Type: EOF, Value: ""})
	return tokens, nil
}
