package main

import (
	"fmt"
	"go_for_fun_projects/helper"
	"log/slog"
)

func start_calculator(logger *slog.Logger) {

	logger.Debug("start calculator - user input requested")

	customer := helper.GetUserInput(Customer)
	investment_amount := helper.ProcessUserInput(InvestInput)
	return_rate := helper.ProcessUserInput(ReturnInput)
	hodle_years := helper.ProcessUserInput(HodlInput)

	logger.Debug("start calculating future values")

	futureValue, futureValueAdjusted := helper.CalculateFutureValue(investment_amount, return_rate, hodle_years)

	helper.SaveFile(&customer, helper.ConvertFloatToByte(futureValueAdjusted))

	fmt.Printf("Your future savings are: %.2f €\n", futureValue)
	fmt.Printf("Your future savings after inflation are: %.2f €\n", futureValueAdjusted)

	logger.Debug("close calculator")
}

func process_user_request(logger *slog.Logger) {
	jsonFile := "user_accounts/user.json"

	firstName := helper.GetUserInput(FirtName)
	lastName := helper.GetUserInput(LastName)
	desiredRequest := helper.GetUserInput(Request)

	user := helper.UserData{
		FirstName: firstName,
		LastName:  lastName,
		Request:   desiredRequest,
	}
	userInput := user.ToString()
	err := user.AppendToJson(jsonFile)
	if err != nil {
		logger.Error("Error saving to Json: %s", "error", err)

	}
	logger.Info("Successfully saved user data to json file", "userInput", userInput)

}

func main() {
	logger := helper.InitializeLogging(DebugMode)

	operation := helper.RequestOperationMode(UserOperationRequest)

	if operation == 1 {
		start_calculator(logger)
	}
	if operation == 2 {
		process_user_request(logger)
	}

}
