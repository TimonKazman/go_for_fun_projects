package helper

import (
	"errors"
	"fmt"
	"math"
	"strconv"
	"strings"
)

func ConvertStringToFloat(string_value string) (float64, error) {
	string_value = strings.Replace(string_value, ",", ".", 1)
	i, err := strconv.ParseFloat(string_value, 32)
	if err != nil {
		return math.NaN(), errors.New("invalide or not a number! Try again")
	}
	if i <= 0.0 {
		return math.NaN(), fmt.Errorf("value cannot be smaller or equal to zero")
	}
	return i, nil
}
func ConvertStringToInt(string_value string) (int64, error) {
	i, err := strconv.ParseInt(string_value, 10, 0)
	if err != nil {
		return math.MinInt64, errors.New("invalide operation! Try again by choosing '1' or '2'")
	}
	return i, nil
}

func ConvertFloatToByte(float_value float64) []byte {
	valueString := fmt.Sprint(float_value)
	return []byte(valueString)
}
