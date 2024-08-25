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
    res, err := evaluateUnitConversion(input)
    if err != nil {
      return "", err
    }

    return fmt.Sprintf("%.2f", res), nil
  }

  return "", fmt.Errorf("Invalid conversion") 
}

func isUnitConversion(input string) bool {
  return strings.Contains(input, "to")
}

func isMathExpression(input string) bool {
  return strings.ContainsAny(input, "+-*/^%")
}
