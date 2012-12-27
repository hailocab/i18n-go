package i18n

import (
	"errors"
)

var (
	ErrLocaleNotFound = errors.New("i18n: locale not found")
)

type Locale struct {
	Code      string
	Language  string
	LanguageISO3 string
	Territory string
	CurrencyCode string
	CurrencySymbol string
	CurrencyDecimalDigits int
	CurrencyDecimalSeparator string
	CurrencyGroupSizes []int
	CurrencyGroupSeparator string
	CurrencyPositivePattern string
	CurrencyNegativePattern string
	ListSeparator string
	NegativeSign string
	PositiveSign string
	PercentSymbol string
	PerMilleSymbol string
	NaNSymbol string
	NumberDecimalDigits int
	NumberDecimalSeparator string
	NumberGroupSizes []int
	NumberGroupSeparator string
	NumberNegativePattern string
}

func GetLocale(code string) (*Locale, error) {
	if locale, found := locales[code]; found {
		return locale, nil
	}
	return nil, ErrLocaleNotFound
}
