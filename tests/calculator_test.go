package tests

import (
	"go_for_fun_projects/helper"
	"reflect"
	"testing"
)

// Test functions need to start with "Test" and need to be in a file with the pattern *_test.go
// Further testing ideas: https://www.youtube.com/watch?v=IqY15Py2N9g&list=WL&index=5

func TestConvertStringToFloat(t *testing.T) {
	number, err := helper.ConvertStringToFloat("100")
	if reflect.TypeOf(number).Kind() != reflect.Float64 || err != nil {
		t.Fatalf("ConvertStringToFloat() = %v, want %v", number, reflect.Float64)
	}
}

func TestConvertFloatToByte(t *testing.T) {
	expected_type := reflect.TypeOf([]byte("200")).Kind()
	conv_number := helper.ConvertFloatToByte(100.34)
	con_number_type := reflect.TypeOf(conv_number).Kind()
	if con_number_type != expected_type {
		t.Fatalf("ConvertFloatToByte() = %v, want %v", con_number_type, expected_type)
	}
}

func TestEndToEnd(t *testing.T) {
	investment_amount := 100_000.0
	return_rate := 8.0
	var hodle_years float64 = 30.0

	futureValue, _ := helper.CalculateFutureValue(investment_amount, return_rate, hodle_years)
	if reflect.TypeOf(futureValue).Kind() != reflect.Float64 {
		t.Fatalf("CalculateFutureValue() = %v, want %v", futureValue, reflect.Float64)
	}
}
