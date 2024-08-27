package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var prevValue float64

func evaluate(input string) (string, error) {
	if strings.HasPrefix(input, "_") {
		if prevValue == 0 {
			return "", fmt.Errorf("You don't have previously saved value")
		}

		// if the input is only "_", show the previously saved value as the result.
		if input == "_" {
			return fmt.Sprintf("%.2f", prevValue), nil
		}

		// otherwise, replace "_" in the input with the previous value.
		val := fmt.Sprintf("%f", prevValue)
		input = strings.Replace(input, "_", val, 1)
	}

	if isMathExpression(input) {
		res, err := evaluateMathExpression(input)
		if err != nil {
			return "", err
		}

		prevValue = math.Round(res*10) / 10

		return fmt.Sprintf("%.2f", res), nil
	}

	if isUnitConversion(input) {
		res, err := evaluateUnitConversion(input)
		if err != nil {
			return "", err
		}

		prevValue = math.Round(res*10) / 10

		return fmt.Sprintf("%.2f", res), nil
	}

	if isRounding(input) {
		res, err := applyRounding(input)
		if err != nil {
			return "", err
		}

		prevValue = math.Round(res*10) / 10

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
