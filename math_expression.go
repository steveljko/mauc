package main

import (
  "fmt"
  "math"
  "strings"
  "strconv"
)

func evaluateMathExpression(input string) (float64, error) {
  tokens, err := tokenize(input)
  if err != nil {
    return 0, err
  }

  postfix, err := toPostfix(tokens)
  if err != nil {
    return 0, err 
  }

  res, err := evaluatePostfix(postfix)
  if err != nil {
    return 0, err
  }

  return res, nil 
}

func priority(token Token) int {
  switch token.Type {
  case PLUS, MINUS:
    return 1
  case MULTIPLY, DIVIDE, MODULO:
    return 2
  case POWER:
    return 3
  default:
    return 0
  }
}

func toPostfix(tokens []Token) (string, error) {
  var output []string
  var operators []Token

  for _, token := range tokens {
    switch token.Type {
    case NUMBER:
      output = append(output, token.Value)
    case PLUS, MINUS, MULTIPLY, DIVIDE, POWER, MODULO:
      for len(operators) > 0 && priority(operators[len(operators)-1]) >= priority(token) {
        output = append(output, operators[len(operators)-1].Value)
        operators = operators[:len(operators)-1]
      }
      operators = append(operators, token)
    case LPAREN:
      operators = append(operators, token)
    case RPAREN:
      for len(operators) > 0 && operators[len(operators)-1].Type != LPAREN {
        output = append(output, operators[len(operators)-1].Value)
        operators = operators[:len(operators)-1]
      }
      if len(operators) == 0 {
        return "", fmt.Errorf("Mismatched parentheses")
      }
      operators = operators[:len(operators)-1]
    case EOF:
      break
    default:
      return "", fmt.Errorf("Unexpected token")
    }
  }

  for len(operators) > 0 {
    if operators[len(operators)-1].Type == LPAREN {
      return "", fmt.Errorf("Mismatched parentheses")
    }
    output = append(output, operators[len(operators)-1].Value)
    operators = operators[:len(operators)-1]
  }

  return strings.Join(output, " "), nil
}

func evaluatePostfix(postfix string) (float64, error) {
  tokens := strings.Fields(postfix)
  stack := []float64{}

  for _, token := range tokens {
    if num, err := strconv.ParseFloat(token, 64); err == nil {
      stack = append(stack, num)
    } else {
      if len(stack) < 2 {
        return 0, fmt.Errorf("Not enough operands for operator")
      }
      b := stack[len(stack)-1]
      a := stack[len(stack)-2]
      stack = stack[:len(stack)-2]

      var result float64
      switch token {
      case "+":
        result = a + b
      case "-":
        result = a - b
      case "*":
        result = a * b
      case "/":
        if b == 0 {
          return 0, fmt.Errorf("Division by zero")
        }
        result = a / b
      case "^":
        result = math.Pow(a, b)
      case "%":
        result = float64(int(a) % int(b))
      default:
        return 0, fmt.Errorf("Unknown operator %s", token)
      }
      stack = append(stack, result)
    }
  }

  if len(stack) != 1 {
    return 0, fmt.Errorf("Invalid postfix expression")
  }

  return stack[0], nil
}
