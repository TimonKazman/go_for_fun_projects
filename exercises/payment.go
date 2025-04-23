package exercises

import "fmt"

type Payment interface {
	ProcessPayment(amount float64) string
}

type SecurePayment interface {
	Payment
	AuthenticateUser() bool
}

type BasePayment struct {
	AccountId     string
	CardNumber    int64
	WalletAddress string
}

func (b BasePayment) ProcessPayment(amount float64) string {

	info := fmt.Sprintf(
		"Processed payment of: $%.2f\nAccount ID: %s\nCard Number: %d\nWallet Adress: %s",
		amount, b.AccountId, b.CardNumber, b.WalletAddress,
	)
	println(info)

	return info
}

type CreditCard struct{ BasePayment }

func (c CreditCard) AuthenticateUser() bool {
	fmt.Println("Logging into credit card account")
	return true
}

type PayPal struct{ BasePayment }

func (p PayPal) AuthenticateUser() bool {
	fmt.Println("Logging into credit card account")
	return true
}

type CryptoWallet struct{ BasePayment }

func (c CryptoWallet) LogIn() {
	fmt.Println("Your in...")
}

//TODO: Use generic types

// Implement a generic function ExecuteTransaction[T Payment](method T, amount float64) that accepts any payment method and calls its ProcessPayment() function.

// Extra Challenge:
// Add validation logic to check if the payment amount is within a reasonable range.

// Allow users to select a payment method dynamically via an input mechanism.
