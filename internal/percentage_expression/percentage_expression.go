package percentage_expression

import (
	"errors"
	"strconv"
	"strings"
)

func Evaluate(input string) (float64, error) {
	parts := strings.Fields(input)

	if len(parts) == 3 {
		percentStr := parts[0]
		op := parts[1]
		numStr := parts[2]

		percent, err := strconv.ParseFloat(strings.TrimSuffix(percentStr, "%"), 64)
		if err != nil {
			return 0, errors.New("invalid percentage value")
		}

		number, err := strconv.ParseFloat(numStr, 64)
		if err != nil {
			return 0, errors.New("invalid number value")
		}

		switch op {
		case "of":
			res := evaluateOf(percent, number)
			return res, nil
		case "off":
			res := evaluateOff(percent, number)
			return res, nil
		}
	}

	return 0, errors.New("can calculate percentage of this expression")
}

// calculates the percentage of a given number
func evaluateOf(percent, number float64) float64 {
	result := (percent / 100) * number
	return result
}

// calculates the discounted price after applying a percentage discount to a given number
func evaluateOff(percent, number float64) float64 {
	discount := (percent / 100) * number
	result := number - discount
	return result
}
