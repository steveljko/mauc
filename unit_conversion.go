package main

import (
  "fmt"
  "strconv"
  "strings"
)

var conversions = map[string]float64 {
  // Length Conversions
  "km_mi": 0.621371,    // kilometers to miles
	"mi_km": 1.60934,     // miles to kilometers
	"m_km":  0.001,       // meters to kilometers
	"km_m":  1000,        // kilometers to meters
	"m_mi":  0.000621371, // meters to miles
	"mi_m":  1609.34,     // miles to meters
	"m_ft":  3.28084,     // meters to feet
	"ft_m":  0.3048,      // feet to meters
	"cm_m":  0.01,        // centimeters to meters
	"m_cm":  100,         // meters to centimeters
	"mm_m":  0.001,       // millimeters to meters
	"m_mm":  1000,        // meters to millimeters

	// Speed Conversions
	"kmh_mph": 0.621371,  // kilometers per hour to miles per hour
	"mph_kmh": 1.60934,   // miles per hour to kilometers per hour

	// Volume Conversions
	"l_ml":  1000,        // liters to milliliters
	"ml_l":  0.001,       // milliliters to liters
	"gal_l": 3.78541,     // gallons to liters
	"l_gal": 0.264172,    // liters to gallons

	// Area Conversions
	"ft2_m2": 0.092903,   // square feet to square meters
	"m2_ft2": 10.7639,    // square meters to square feet
	"ac_m2":  4046.86,    // acres to square meters
	"m2_ac":  0.000247105, // square meters to acres
	"ha_m2":  1e4,        // hectares to square meters
	"m2_ha":  0.0001,     // square meters to hectares

	// Weight Conversions
	"kg_lb":  2.20462,    // kilograms to pounds
	"lb_kg":  0.453592,   // pounds to kilograms
	"g_kg":   0.001,      // grams to kilograms
	"kg_g":   1000,       // kilograms to grams
	"oz_lb":  0.0625,     // ounces to pounds
	"lb_oz":  16,         // pounds to ounces
}

func evaluateUnitConversion(input string) (float64, error) {
  parts := strings.Split(input, " ")

  if len(parts) == 4 {
    value, err := strconv.ParseFloat(parts[0], 64)
		if err != nil {
			return 0, fmt.Errorf("Invalid conversion value")
		}

    from := strings.ToLower(parts[1])
    to := strings.ToLower(parts[3])
    key := from + "_" + to

    factor, ok := conversions[key]
    if !ok {
      return 0, fmt.Errorf("This conversion is not allowed")
    }

    res := value * factor

    return res, nil
  }

  return 0, fmt.Errorf("Invalid unit conversion")
}
