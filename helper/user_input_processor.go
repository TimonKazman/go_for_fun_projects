package helper

import (
	"bufio"
	"encoding/json"
	"fmt"
	"math"
	"os"
	"strings"
)

const InflationRatePercent = 2.4

type UserData struct {
	FirstName string
	LastName  string
	Request   string
}

// Method
func (u *UserData) ToString() string {
	return fmt.Sprintf("%s;%s;%s", u.FirstName, u.LastName, u.Request)
}

func (u *UserData) AppendToJson(filename string) error {
	var users []UserData

	if _, err := os.Stat(filename); err == nil {
		data, err := os.ReadFile(filename)
		if err != nil {
			return fmt.Errorf("failed to read file: %v", err)
		}
		if err := json.Unmarshal(data, &users); err != nil {
			return fmt.Errorf("failed to unmarshal JSON: %v", err)
		}
	}

	users = append(users, *u)

	updatedData, err := json.Marshal(users)
	if err != nil {
		return fmt.Errorf("failed to marshal JSON: %v", err)
	}
	return os.WriteFile(filename, updatedData, 0644)
}

func GetUserInput(input_request string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(input_request)
	input, _ := reader.ReadString(('\n'))
	return strings.TrimSpace((input))
}

func ProcessUserInput(inputString string) float64 {
	var userInputNumber float64
	for {
		userInputString := GetUserInput(inputString)
		var err error
		userInputNumber, err = ConvertStringToFloat(userInputString)
		if err == nil {
			return userInputNumber
		}
		fmt.Println(err)
	}
}

func RequestOperationMode(inputString string) int64 {
	var userInputNumber int64
	for {
		userInputString := GetUserInput(inputString)
		var err error
		userInputNumber, err = ConvertStringToInt(userInputString)
		if err == nil && (userInputNumber == 1 || userInputNumber == 2) {
			return userInputNumber
		}
		fmt.Println("invalide operation! Try again by choosing '1' or '2'.")
	}
}

func CalculateFutureValue(investmentAmount, returnRate, hodleYears float64) (float64, float64) {
	/* Calculates the savings based on (investAmount * returnrate^years) and fixed value after Inflation*/
	var hundredPercent float64 = 100

	futureValue := investmentAmount * math.Pow(1+(returnRate/hundredPercent), hodleYears)
	futureValueAdjusted := futureValue * math.Pow(1-(InflationRatePercent/hundredPercent), hodleYears)

	return futureValue, futureValueAdjusted

}
