package currency

// Currency represets all details about a currency.
type Currency struct {
	Code string
	// Symbol is the symbol used in the locale, e.g. € for Euro.
	Symbol string
	// DecimalDigits is the number of digits after the decimal point.
	DecimalDigits int
	// DecimalSeparator is the string used to separate the currency value.
	DecimalSeparator string
	// GroupSizes is the size to be used for grouping a currency value,
	// e.g. 3 for currencies to be formatted like 123.456.789,01
	GroupSizes []int
	// GroupSeparator is the seperator used for currency grouping.
	GroupSeparator string
	// PositivePattern is the pattern used for positive currency values.
	PositivePattern string
	// NegativePattern is the pattern used for negative currency values.
	NegativePattern string
}

func Get(code string) *Currency {
	return currencies[code]
}

func Currencies() map[string]*Currency {
	return currencies
}
