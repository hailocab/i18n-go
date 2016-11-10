# i18n-go

This is an internationalization library forked from [github.com/olivere/i18n-go](https://github.com/olivere/i18n-go)

## Example Usage
	package main

	import (
		"fmt"
		"github.com/River-Island/i18n-go/money"
	)

	func main() {
		// Initialise new Money
		m := money.New(1050, "USD")

		// Or Initialise using the main unit
		m = money.NewFromMainUnit("10.50", "USD")

		// Get the value of Money
		fmt.Println(m.Value()) // 1050

		// Get the float64 super unit value
		fmt.Println(m.Get()) // 10.5

		// Get a rounded int64 value
		fmt.Println(m.Gett()) // 10

		// String value
		fmt.Println(m.String()) // 10.50 USD

		// Format
		fmt.Println(m.Format("en_US")) // $10.50

		// More money
		n := money.New(1000, "USD")

		// Add
		m.Add(n) // 2050

		// Subtract
		m.Sub(n) // 1050

		// Multiply
		m.Mul(n) // 10500

		// Chain operations
		money.New(1050, "USD").Add(n).Sub(money.New(200, "USD")).Get() // 18.5

		// Set the decimal place
		money.SetDecimal(4)

		// Set decimal by currency
		money.SetDecimalByCurrency("USD")

		// Set decimal by locale
		money.SetDecimalByLocale("en_US")
	}
