package money

import (
	"bytes"
	"fmt"
	"math"
	"strings"

	"github.com/hailocab/i18n-go/currency"
	"github.com/hailocab/i18n-go/locale"
)

const (
	defaultSeparator       = "."
	defaultGroupSeparator  = ""
	defaultGroupSize       = 3
	defaultDecimalDigits   = 2
	defaultPositivePattern = "n $"
	defaultNegativePattern = "-n $"
)

type MoneyFormatter struct {
	decimalSeparator string
	groupSeparator   string
	groupSize        int
	decimalDigits    int
	positivePattern  string
	negativePattern  string
	currencySymbol   string
}

// GetFormatter returns a configured formatter given the locale and currency
// uses defaults for certain fields if locale or currency are blank/not recognised
func GetFormatter(locString, currencyCode string) *MoneyFormatter {
	formatter := &MoneyFormatter{
		decimalSeparator: defaultSeparator,
		groupSeparator:   defaultGroupSeparator,
		groupSize:        defaultGroupSize,
		decimalDigits:    defaultDecimalDigits,
		positivePattern:  defaultPositivePattern,
		negativePattern:  defaultNegativePattern,
		currencySymbol:   currencyCode,
	}
	if curr := currency.Get(currencyCode); curr != nil {
		formatter.currencySymbol = curr.Symbol
		formatter.decimalDigits = curr.DecimalDigits
	}
	if loc := locale.Get(locString); loc != nil {
		formatter.decimalSeparator = loc.CurrencyDecimalSeparator
		formatter.groupSeparator = loc.CurrencyGroupSeparator
		if len(loc.CurrencyGroupSizes) >= 1 {
			formatter.groupSize = loc.CurrencyGroupSizes[0]
		}
		formatter.negativePattern = loc.CurrencyNegativePattern
		formatter.positivePattern = loc.CurrencyPositivePattern
	} else {
		formatter.currencySymbol = currencyCode // do this because default format pattern puts the currency code at the end.
	}
	return formatter
}

func (mf *MoneyFormatter) Format(value int64) string {
	// DP is a measure for decimals: 2 decimal digits => dp = 10^2
	dp := int64(math.Pow10(mf.decimalDigits))

	// Group DP is a measure for grouping: 3 decimal digits => groupDp = 10^3
	groupDp := int64(math.Pow10(mf.groupSize))

	// We use absolute values (as int64) from here on, because the
	// negative sign is part of the currency format pattern.
	absVal := value
	if value < 0 {
		absVal = -absVal
	}
	wholeVal := absVal / dp
	decVal := absVal % dp

	// The unformatted string (without grouping and with a decimal sep of ".")
	var unformatted string
	if mf.decimalDigits > 0 {
		unformatted = fmt.Sprintf("%d.%0"+fmt.Sprintf("%d", mf.decimalDigits)+"d", wholeVal, decVal)
	} else {
		unformatted = fmt.Sprintf("%d", wholeVal)
	}

	// Perform grouping operation of the whole number
	groups := make([]string, 0)
	innerGroupFmt := "%0" + fmt.Sprintf("%d", mf.groupSize) + "d"
	for {
		group := wholeVal % groupDp
		var s string
		if wholeVal < groupDp {
			s = fmt.Sprintf("%d", group)
		} else {
			s = fmt.Sprintf(innerGroupFmt, group)
		}
		groups = append(groups, s)
		wholeVal /= groupDp
		if wholeVal == 0 {
			break
		}
	}
	var wholeBuf bytes.Buffer
	for i, _ := range groups {
		if i > 0 {
			wholeBuf.WriteString(mf.groupSeparator)
		}
		wholeBuf.WriteString(groups[len(groups)-i-1])
	}

	// Which pattern do we need?
	// Notice that the minus sign is part of the pattern
	var pattern string
	if value >= 0 {
		pattern = mf.positivePattern
	} else {
		pattern = mf.negativePattern
	}

	// Split into whole and decimal and build formatted number
	var formatted string
	parts := strings.SplitN(unformatted, ".", 2)
	if len(parts) > 1 {
		formatted = fmt.Sprintf("%s%s%s", wholeBuf.String(), mf.decimalSeparator, parts[1])
	} else {
		formatted = wholeBuf.String()
	}
	output := strings.Replace(pattern, "$", mf.currencySymbol, -1)
	output = strings.Replace(output, "n", formatted, -1)

	return output
}
