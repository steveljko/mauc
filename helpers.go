package main

import (
  "fmt"
  "strings"
)

func evaluate(input string) (string, error) {
  if isMathExpression(input) {
    return input, nil
  }

  if isUnitConversion(input) {
    return input, nil
  }

  return nil, fmt.Errorf("Invalid conversion") 
}

func isUnitConversion(input string) bool {
  return strings.Contains(input, "to")
}

func isMathExpression(input string) bool {
  return strings.ContainsAny(input, "+-*/^%")
}
