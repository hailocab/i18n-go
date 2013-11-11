package generator

import (
	"fmt"
	"github.com/hailocab/i18n-go/currency"
	"github.com/hailocab/i18n-go/locale"
	"os"
	"sort"
)

func Currency() {
	currencies := make(map[string]*currency.Currency)
	for _, v := range locale.Locales() {
		c := &Currency{
			Code:             v.CurrencyCode,
			Symbol:           v.CurrencySymbol,
			DecimalDigits:    v.CurrencyDecimalDigits,
			DecimalSeparator: v.CurrencyDecimalSeparator,
			GroupSizes:       v.CurrencyGroupSizes,
			GroupSeparator:   v.CurrencyGroupSeparator,
			PositivePattern:  v.CurrencyPositivePattern,
			NegativePattern:  v.CurrencyNegativePattern,
		}
		currencies[v.CurrencyCode] = c
	}

	keys := []string{}
	for k, _ := range currencies {
		keys = append(keys, k)
	}
	sort.Strings(keys)

	os.Stdout.Write([]byte("package currency\n\n"))
	os.Stdout.Write([]byte("var currencies = map[string]*Currency{\n"))
	outputString := "\t\t%s: \"%v\",\n"
	for _, k := range keys {
		os.Stdout.Write([]byte(fmt.Sprintf("\t\"%s\": &Currency{\n", k)))
		os.Stdout.Write([]byte(fmt.Sprintf(outputString, "Code", currencies[k].Code)))
		os.Stdout.Write([]byte(fmt.Sprintf(outputString, "Symbol", currencies[k].Symbol)))
		os.Stdout.Write([]byte(fmt.Sprintf("\t\t%s: %v,\n", "DecimalDigits", currencies[k].DecimalDigits)))
		os.Stdout.Write([]byte(fmt.Sprintf(outputString, "DecimalSeparator", currencies[k].DecimalSeparator)))
		os.Stdout.Write([]byte(fmt.Sprintf("\t\t%s: %#v,\n", "GroupSizes", currencies[k].GroupSizes)))
		os.Stdout.Write([]byte(fmt.Sprintf(outputString, "GroupSeparator", currencies[k].GroupSeparator)))
		os.Stdout.Write([]byte(fmt.Sprintf(outputString, "PositivePattern", currencies[k].PositivePattern)))
		os.Stdout.Write([]byte(fmt.Sprintf(outputString, "NegativePattern", currencies[k].NegativePattern)))
		os.Stdout.Write([]byte("\t},\n"))
	}
	os.Stdout.Write([]byte("}\n"))
}
