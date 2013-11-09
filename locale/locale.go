package locale

// Locale contains all information about a locale.
type Locale struct {
	// Code is the ISO code of the locale in the xx_YY format, e.g. de_AT.
	Code string
	// Language is the 2-letter, downcase ISO code of the language, e.g. de.
	Language string
	// LanguageISO3 is the 3-letter, downcase ISO code of the language, e.g. deu.
	LanguageISO3 string
	// Territory is the 2-letter, upcase ISO code of the territory, e.g. DE.
	Territory string
	// CurrencyCode is the 3-letter, upcase ISO code of the currency
	// in the locale, e.g. EUR for Germany.
	CurrencyCode string
	// CurrencySymbol is the symbol used in the locale, e.g. € for Euro.
	CurrencySymbol string
	// CurrencyDecimalDigits is the number of digits after the decimal point.
	CurrencyDecimalDigits int
	// CurrencyDecimalSeparator is the string used to separate the currency value.
	CurrencyDecimalSeparator string
	// CurrencyGroupSizes is the size to be used for grouping a currency value,
	// e.g. 3 for currencies to be formatted like 123.456.789,01
	CurrencyGroupSizes []int
	// CurrencyGroupSeparator is the seperator used for currency grouping.
	CurrencyGroupSeparator string
	// CurrencyPositivePattern is the pattern used for positive currency values.
	CurrencyPositivePattern string
	// CurrencyNegativePattern is the pattern used for negative currency values.
	CurrencyNegativePattern string
	// ListSeparator is the seperator used for lists, e.g. a comma.
	ListSeparator string
	// NegativeSign is the symbol to be used for negative numeric values.
	NegativeSign string
	// PositiveSign is the symbol to be used for positive numeric values.
	PositiveSign string
	// PercentSymbol is the symbol to be used for percentages.
	PercentSymbol string
	// PerMilleSymbol is the symbol to be used for per-mille values.
	PerMilleSymbol string
	// NaNSymbol is the symbol to be used for non-numbers.
	NaNSymbol string
	// NumberDecimalDigits is the number of digits after the decimal point.
	NumberDecimalDigits int
	// NumberDecimalSeparator is the string used to separate the currency value.
	NumberDecimalSeparator string
	// NumberGroupSizes is the size to be used for grouping a currency value,
	// e.g. 3 for currencies to be formatted like 123.456.789,01
	NumberGroupSizes []int
	// NumberGroupSeparator is the seperator used for currency grouping.
	NumberGroupSeparator string
	// NumberNegativePattern is the pattern used for negative currency values.
	NumberNegativePattern string
}

func Get(code string) *Locale {
	return locales[code]
}

func Locales() map[string]*Locale {
	return locales
}
