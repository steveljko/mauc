package main

import (
  "fmt"
  "math"
  "strconv"
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

  if isRounding(input) {
    res, err := applyRounding(input)
    if err != nil {
      return "", err
    }

    return fmt.Sprintf("%.0f", res), nil
  }

  return "", fmt.Errorf("Invalid conversion") 
}

func isUnitConversion(input string) bool {
  return strings.Contains(input, "to")
}

func isMathExpression(input string) bool {
  return strings.ContainsAny(input, "+-*/^%")
}

func isRounding(input string) bool {
  return strings.HasPrefix(input, "round") || strings.HasPrefix(input, "r")
}

func applyRounding(input string) (float64, error) {
  parts := strings.Split(input, " ")

  value, err := strconv.ParseFloat(parts[1], 64)
  if err != nil {
    return 0, fmt.Errorf("Invalid number for rounding")
  }

  return math.Round(value), nil
}
