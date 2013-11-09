package currency

// Currency represets all details about a currency.
type Currency struct {
	// Code is the 3-letter ISO code of the currency
	Code string
	// Symbol is the common symbol used for the currency, e.g. € for Euro.
	Symbol string
}

func Get(code string) *Currency {
	return currencies[code]
}

func Currencies() map[string]*Currency {
	return currencies
}
