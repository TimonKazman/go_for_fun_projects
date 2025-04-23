package main

import (
	"fmt"
	"go_for_fun_projects/helper"
	"log"
	"log/slog"

	"github.com/manifoldco/promptui"
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

func select_payment(logger *slog.Logger) {
	items := []string{"Credit Card", "PayPal", "Crypto Wallet"}

	prompt := promptui.Select{
		Label: "Select Payment Method",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	switch result {
	case "Credit Card":
		//TODO:
	case "PayPal":
		//TODO:
	case "Crypto Wallet":
		//TODO:
	}
}

func main() {
	logger := helper.InitializeLogging(DebugMode)
	//TODO: create a generic terminal interaction function
	items := []string{"Investment Calculator", "Create User Account", "Payment"}

	prompt := promptui.Select{
		Label: "Select Operation Mode",
		Items: items,
	}

	_, result, err := prompt.Run()
	if err != nil {
		log.Fatalf("Prompt failed %v\n", err)
	}

	switch result {
	case "Investment Calculator":
		start_calculator(logger)
	case "Create User Account":
		process_user_request(logger)
	case "Payment":
		select_payment(logger)
	}
}
