package money

import (
	"fmt"
	"reflect"
	"testing"
)

func TestM(t *testing.T) {
	m := &Money{}
	if m.M != 0 {
		t.Errorf("expected money amount to be %v, got %v", 0, m.M)
	}

	m = &Money{-123, "EUR"}
	if m.M != -123 {
		t.Errorf("expected money amount to be %v, got %v", -123, m.M)
	}

	m = &Money{123, "EUR"}
	if m.M != 123 {
		t.Errorf("expected money amount to be %v, got %v", 123, m.M)
	}
}

func TestGet(t *testing.T) {
	m := &Money{123, "EUR"}
	if m.Get() != 1.23 {
		t.Errorf("expected money amount to be %v, got %v", 1.23, m.Get())
	}

	m = &Money{12345, "EUR"}
	if m.Get() != 123.45 {
		t.Errorf("expected money amount to be %v, got %v", 123.45, m.Get())
	}
}

func TestSet(t *testing.T) {
	m := &Money{123, "EUR"}
	if m.Get() != 1.23 {
		t.Errorf("expected money amount to be %v, got %v", 1.23, m.Get())
	}
	m.Set(12345)
	if m.Get() != 123.45 {
		t.Errorf("expected money amount to be %v, got %v", 123.45, m.Get())
	}
}

func TestSetf(t *testing.T) {
	m := &Money{123, "EUR"}
	if m.Get() != 1.23 {
		t.Errorf("expected money amount to be %v, got %v", 1.23, m.Get())
	}
	m.Set(12345)
	if m.Get() != 123.45 {
		t.Errorf("expected money amount to be %v, got %v", 123.45, m.Get())
	}
}

func TestAdd(t *testing.T) {
	m1 := &Money{123, "EUR"}
	m2 := &Money{678, "EUR"}
	m3 := m1.Add(m2)
	if m3.Get() != 8.01 {
		t.Errorf("expected money amount to be %v, got %v", 8.01, m3.Get())
	}
}

func TestMul(t *testing.T) {
	m1 := &Money{123, "EUR"}
	m2 := &Money{200, "EUR"}
	m3 := m1.Mul(m2)
	if m3.Get() != 2.46 {
		t.Errorf("expected money amount to be %v, got %v", 2.46, m3.Get())
	}
}

func TestMulf(t *testing.T) {
	m1 := &Money{123, "EUR"}
	m2 := 2.0
	m3 := m1.Mulf(m2)
	if m3.Get() != 2.46 {
		t.Errorf("expected money amount to be %v, got %v", 2.46, m3.Get())
	}
}

func TestMoneyStringer(t *testing.T) {
	var fixtures = []struct {
		m        *Money
		expected string
	}{
		{&Money{0, "EUR"}, "0.00 EUR"},
		{&Money{1, "EUR"}, "0.01 EUR"},
		{&Money{12, "EUR"}, "0.12 EUR"},
		{&Money{123, "EUR"}, "1.23 EUR"},
		{&Money{1234, "EUR"}, "12.34 EUR"},
		{&Money{123456, "EUR"}, "1234.56 EUR"},
		{&Money{1234567, "EUR"}, "12345.67 EUR"},
		{&Money{1234567890, "EUR"}, "12345678.90 EUR"},
		{&Money{-1, "EUR"}, "-0.01 EUR"},
		{&Money{-12, "EUR"}, "-0.12 EUR"},
		{&Money{-123, "EUR"}, "-1.23 EUR"},
		{&Money{-1234, "EUR"}, "-12.34 EUR"},
		{&Money{-123456, "EUR"}, "-1234.56 EUR"},
		{&Money{-1234567, "EUR"}, "-12345.67 EUR"},
		{&Money{-1234567890, "EUR"}, "-12345678.90 EUR"},
		{&Money{0, "GBP"}, "0.00 GBP"},
		{&Money{1, "GBP"}, "0.01 GBP"},
		{&Money{12, "GBP"}, "0.12 GBP"},
		{&Money{123, "GBP"}, "1.23 GBP"},
		{&Money{1234, "GBP"}, "12.34 GBP"},
		{&Money{123456, "GBP"}, "1234.56 GBP"},
		{&Money{1234567, "GBP"}, "12345.67 GBP"},
		{&Money{1234567890, "GBP"}, "12345678.90 GBP"},
		{&Money{-1, "GBP"}, "-0.01 GBP"},
		{&Money{-12, "GBP"}, "-0.12 GBP"},
		{&Money{-123, "GBP"}, "-1.23 GBP"},
		{&Money{-1234, "GBP"}, "-12.34 GBP"},
		{&Money{-123456, "GBP"}, "-1234.56 GBP"},
		{&Money{-1234567, "GBP"}, "-12345.67 GBP"},
		{&Money{-1234567890, "GBP"}, "-12345678.90 GBP"},
		{&Money{0, "JPY"}, "0 JPY"},
		{&Money{1, "JPY"}, "1 JPY"},
		{&Money{12, "JPY"}, "12 JPY"},
		{&Money{123, "JPY"}, "123 JPY"},
		{&Money{1234, "JPY"}, "1234 JPY"},
		{&Money{123456, "JPY"}, "123456 JPY"},
		{&Money{1234567, "JPY"}, "1234567 JPY"},
		{&Money{1234567890, "JPY"}, "1234567890 JPY"},
		{&Money{-1, "JPY"}, "-1 JPY"},
		{&Money{-12, "JPY"}, "-12 JPY"},
		{&Money{-123, "JPY"}, "-123 JPY"},
		{&Money{-1234, "JPY"}, "-1234 JPY"},
		{&Money{-123456, "JPY"}, "-123456 JPY"},
		{&Money{-1234567, "JPY"}, "-1234567 JPY"},
		{&Money{-1234567890, "JPY"}, "-1234567890 JPY"},
	}

	for _, f := range fixtures {
		got := fmt.Sprintf("%s", f.m)
		if got != f.expected {
			t.Errorf("expected %s, got %s", f.expected, got)
		}
	}
}

func TestMoneyFormat(t *testing.T) {
	var fixtures = []struct {
		m        *Money
		locale   string
		expected string
	}{
		{&Money{0, "EUR"}, "de", "0.00 EUR"},
		{&Money{1, "EUR"}, "de", "0.01 EUR"},
		{&Money{12, "EUR"}, "de", "0.12 EUR"},
		{&Money{123, "EUR"}, "de", "1.23 EUR"},
		{&Money{1234, "EUR"}, "de", "12.34 EUR"},
		{&Money{123456, "EUR"}, "de", "1234.56 EUR"},
		{&Money{1234567, "EUR"}, "de", "12345.67 EUR"},
		{&Money{1234567890, "EUR"}, "de", "12345678.90 EUR"},
		{&Money{-1, "EUR"}, "de", "-0.01 EUR"},
		{&Money{-12, "EUR"}, "de", "-0.12 EUR"},
		{&Money{-123, "EUR"}, "de", "-1.23 EUR"},
		{&Money{-1234, "EUR"}, "de", "-12.34 EUR"},
		{&Money{-123456, "EUR"}, "de", "-1234.56 EUR"},
		{&Money{-1234567, "EUR"}, "de", "-12345.67 EUR"},
		{&Money{-1234567890, "EUR"}, "de", "-12345678.90 EUR"},

		{&Money{0, "EUR"}, "de_DE", "0,00 €"},
		{&Money{1, "EUR"}, "de_DE", "0,01 €"},
		{&Money{12, "EUR"}, "de_DE", "0,12 €"},
		{&Money{123, "EUR"}, "de_DE", "1,23 €"},
		{&Money{1234, "EUR"}, "de_DE", "12,34 €"},
		{&Money{123456, "EUR"}, "de_DE", "1.234,56 €"},
		{&Money{1234567, "EUR"}, "de_DE", "12.345,67 €"},
		{&Money{1234567890, "EUR"}, "de_DE", "12.345.678,90 €"},
		{&Money{-1, "EUR"}, "de_DE", "-0,01 €"},
		{&Money{-12, "EUR"}, "de_DE", "-0,12 €"},
		{&Money{-123, "EUR"}, "de_DE", "-1,23 €"},
		{&Money{-1234, "EUR"}, "de_DE", "-12,34 €"},
		{&Money{-123456, "EUR"}, "de_DE", "-1.234,56 €"},
		{&Money{-1234567, "EUR"}, "de_DE", "-12.345,67 €"},
		{&Money{-1234567890, "EUR"}, "de_DE", "-12.345.678,90 €"},

		{&Money{0, "EUR"}, "de_AT", "€ 0,00"},
		{&Money{1, "EUR"}, "de_AT", "€ 0,01"},
		{&Money{12, "EUR"}, "de_AT", "€ 0,12"},
		{&Money{123, "EUR"}, "de_AT", "€ 1,23"},
		{&Money{1234, "EUR"}, "de_AT", "€ 12,34"},
		{&Money{123456, "EUR"}, "de_AT", "€ 1.234,56"},
		{&Money{1234567, "EUR"}, "de_AT", "€ 12.345,67"},
		{&Money{1234567890, "EUR"}, "de_AT", "€ 12.345.678,90"},
		{&Money{-1, "EUR"}, "de_AT", "-€ 0,01"},
		{&Money{-12, "EUR"}, "de_AT", "-€ 0,12"},
		{&Money{-123, "EUR"}, "de_AT", "-€ 1,23"},
		{&Money{-1234, "EUR"}, "de_AT", "-€ 12,34"},
		{&Money{-123456, "EUR"}, "de_AT", "-€ 1.234,56"},
		{&Money{-1234567, "EUR"}, "de_AT", "-€ 12.345,67"},
		{&Money{-1234567890, "EUR"}, "de_AT", "-€ 12.345.678,90"},

		{&Money{0, "EUR"}, "de_CH", "€ 0.00"},
		{&Money{1, "EUR"}, "de_CH", "€ 0.01"},
		{&Money{12, "EUR"}, "de_CH", "€ 0.12"},
		{&Money{123, "EUR"}, "de_CH", "€ 1.23"},
		{&Money{1234, "EUR"}, "de_CH", "€ 12.34"},
		{&Money{123456, "EUR"}, "de_CH", "€ 1'234.56"},
		{&Money{1234567, "EUR"}, "de_CH", "€ 12'345.67"},
		{&Money{1234567890, "EUR"}, "de_CH", "€ 12'345'678.90"},
		{&Money{-1, "EUR"}, "de_CH", "€-0.01"},
		{&Money{-12, "EUR"}, "de_CH", "€-0.12"},
		{&Money{-123, "EUR"}, "de_CH", "€-1.23"},
		{&Money{-1234, "EUR"}, "de_CH", "€-12.34"},
		{&Money{-123456, "EUR"}, "de_CH", "€-1'234.56"},
		{&Money{-1234567, "EUR"}, "de_CH", "€-12'345.67"},
		{&Money{-1234567890, "EUR"}, "de_CH", "€-12'345'678.90"},

		{&Money{0, "EUR"}, "en", "0.00 EUR"},
		{&Money{1, "EUR"}, "en", "0.01 EUR"},
		{&Money{12, "EUR"}, "en", "0.12 EUR"},
		{&Money{123, "EUR"}, "en", "1.23 EUR"},
		{&Money{1234, "EUR"}, "en", "12.34 EUR"},
		{&Money{123456, "EUR"}, "en", "1234.56 EUR"},
		{&Money{1234567, "EUR"}, "en", "12345.67 EUR"},
		{&Money{1234567890, "EUR"}, "en", "12345678.90 EUR"},
		{&Money{-1, "EUR"}, "en", "-0.01 EUR"},
		{&Money{-12, "EUR"}, "en", "-0.12 EUR"},
		{&Money{-123, "EUR"}, "en", "-1.23 EUR"},
		{&Money{-1234, "EUR"}, "en", "-12.34 EUR"},
		{&Money{-123456, "EUR"}, "en", "-1234.56 EUR"},
		{&Money{-1234567, "EUR"}, "en", "-12345.67 EUR"},
		{&Money{-1234567890, "EUR"}, "en", "-12345678.90 EUR"},

		{&Money{0, "EUR"}, "en_US", "€0.00"},
		{&Money{1, "EUR"}, "en_US", "€0.01"},
		{&Money{12, "EUR"}, "en_US", "€0.12"},
		{&Money{123, "EUR"}, "en_US", "€1.23"},
		{&Money{1234, "EUR"}, "en_US", "€12.34"},
		{&Money{123456, "EUR"}, "en_US", "€1,234.56"},
		{&Money{1234567, "EUR"}, "en_US", "€12,345.67"},
		{&Money{1234567890, "EUR"}, "en_US", "€12,345,678.90"},
		{&Money{-1, "EUR"}, "en_US", "(€0.01)"},
		{&Money{-12, "EUR"}, "en_US", "(€0.12)"},
		{&Money{-123, "EUR"}, "en_US", "(€1.23)"},
		{&Money{-1234, "EUR"}, "en_US", "(€12.34)"},
		{&Money{-123456, "EUR"}, "en_US", "(€1,234.56)"},
		{&Money{-1234567, "EUR"}, "en_US", "(€12,345.67)"},
		{&Money{-1234567890, "EUR"}, "en_US", "(€12,345,678.90)"},

		{&Money{0, "EUR"}, "fr", "0.00 EUR"},
		{&Money{1, "EUR"}, "fr", "0.01 EUR"},
		{&Money{12, "EUR"}, "fr", "0.12 EUR"},
		{&Money{123, "EUR"}, "fr", "1.23 EUR"},
		{&Money{1234, "EUR"}, "fr", "12.34 EUR"},
		{&Money{123456, "EUR"}, "fr", "1234.56 EUR"},
		{&Money{1234567, "EUR"}, "fr", "12345.67 EUR"},
		{&Money{1234567890, "EUR"}, "fr", "12345678.90 EUR"},
		{&Money{-1, "EUR"}, "fr", "-0.01 EUR"},
		{&Money{-12, "EUR"}, "fr", "-0.12 EUR"},
		{&Money{-123, "EUR"}, "fr", "-1.23 EUR"},
		{&Money{-1234, "EUR"}, "fr", "-12.34 EUR"},
		{&Money{-123456, "EUR"}, "fr", "-1234.56 EUR"},
		{&Money{-1234567, "EUR"}, "fr", "-12345.67 EUR"},
		{&Money{-1234567890, "EUR"}, "fr", "-12345678.90 EUR"},

		{&Money{1234567890, "USD"}, "en_US", "$12,345,678.90"},
		{&Money{-1234567890, "USD"}, "en_US", "($12,345,678.90)"},
		{&Money{1234567890, "USD"}, "de_DE", "12.345.678,90 $"},
		{&Money{-1234567890, "USD"}, "de_DE", "-12.345.678,90 $"},
		{&Money{1234567890, "USD"}, "de_CH", "$ 12'345'678.90"},
		{&Money{-1234567890, "USD"}, "de_CH", "$-12'345'678.90"},
		{&Money{1234567890, "GBP"}, "en_GB", "£12,345,678.90"},
		{&Money{-1234567890, "GBP"}, "en_GB", "-£12,345,678.90"},
		{&Money{1234567890, "GBP"}, "de_DE", "12.345.678,90 £"},
		{&Money{-1234567890, "GBP"}, "de_DE", "-12.345.678,90 £"},
		{&Money{1234567890, "GBP"}, "de_CH", "£ 12'345'678.90"},
		{&Money{-1234567890, "GBP"}, "de_CH", "£-12'345'678.90"},
		{&Money{1234567890, "HUF"}, "hu_HU", "12 345 678,90 Ft"},
		{&Money{-1234567890, "HUF"}, "hu_HU", "-12 345 678,90 Ft"},
		{&Money{1234567890, "HUF"}, "de_DE", "12.345.678,90 Ft"},
		{&Money{-1234567890, "HUF"}, "de_DE", "-12.345.678,90 Ft"},
		{&Money{1234567890, "HUF"}, "de_CH", "Ft 12'345'678.90"},
		{&Money{-1234567890, "HUF"}, "de_CH", "Ft-12'345'678.90"},
		{&Money{1234567890, "JPY"}, "ja_JP", "¥1,234,567,890"},
		{&Money{-1234567890, "JPY"}, "ja_JP", "-¥1,234,567,890"},
		{&Money{1234567890, "JPY"}, "de_DE", "1.234.567.890 ¥"},
		{&Money{-1234567890, "JPY"}, "de_DE", "-1.234.567.890 ¥"},
		{&Money{1234567890, "JPY"}, "de_CH", "¥ 1'234'567'890"},
		{&Money{-1234567890, "JPY"}, "de_CH", "¥-1'234'567'890"},
		{&Money{1234567890, "JPY"}, "en_GB", "¥1,234,567,890"},
		{&Money{10, "JPY"}, "ja_JP", "¥10"},
		{&Money{1200, "JPY"}, "ja_JP", "¥1,200"},
		{&Money{10, "JPY"}, "en_JP", "10 JPY"},
		{&Money{1200, "JPY"}, "en_JP", "1200 JPY"},
		{&Money{1234567890, "SEK"}, "se_SE", "12.345.678,90 kr"},
		{&Money{-1234567890, "SEK"}, "se_SE", "-12.345.678,90 kr"},
		{&Money{1234567890, "SEK"}, "de_DE", "12.345.678,90 kr"},
		{&Money{-1234567890, "SEK"}, "de_DE", "-12.345.678,90 kr"},
		{&Money{1234567890, "SEK"}, "de_CH", "kr 12'345'678.90"},
		{&Money{-1234567890, "SEK"}, "de_CH", "kr-12'345'678.90"},

		{&Money{200000, "ZAR"}, "en_ZA", "R 2 000,00"},
		{&Money{120000, "ZAR"}, "en_ZA", "R 1 200,00"},
		{&Money{90000, "ZAR"}, "en_ZA", "R 900,00"},
		{&Money{90001, "ZAR"}, "en_ZA", "R 900,01"},
	}

	for _, f := range fixtures {
		got := f.m.Format(f.locale)
		if got != f.expected {
			t.Errorf("expected %s, got %s (currency: %s, locale: %s)", f.expected, got, f.m.C, f.locale)
		}
	}
}

func TestMoneyFormatNoSymbol(t *testing.T) {
	var fixtures = []struct {
		m        *Money
		locale   string
		expected string
	}{
		{&Money{0, "EUR"}, "de", "0.00"},
		{&Money{1, "EUR"}, "de", "0.01"},
		{&Money{12, "EUR"}, "de", "0.12"},
		{&Money{123, "EUR"}, "de", "1.23"},
		{&Money{1234, "EUR"}, "de", "12.34"},
		{&Money{123456, "EUR"}, "de", "1234.56"},
		{&Money{1234567, "EUR"}, "de", "12345.67"},
		{&Money{1234567890, "EUR"}, "de", "12345678.90"},
		{&Money{-1, "EUR"}, "de", "-0.01"},
		{&Money{-12, "EUR"}, "de", "-0.12"},
		{&Money{-123, "EUR"}, "de", "-1.23"},
		{&Money{-1234, "EUR"}, "de", "-12.34"},
		{&Money{-123456, "EUR"}, "de", "-1234.56"},
		{&Money{-1234567, "EUR"}, "de", "-12345.67"},
		{&Money{-1234567890, "EUR"}, "de", "-12345678.90"},

		{&Money{0, "EUR"}, "de_DE", "0,00"},
		{&Money{1, "EUR"}, "de_DE", "0,01"},
		{&Money{12, "EUR"}, "de_DE", "0,12"},
		{&Money{123, "EUR"}, "de_DE", "1,23"},
		{&Money{1234, "EUR"}, "de_DE", "12,34"},
		{&Money{123456, "EUR"}, "de_DE", "1.234,56"},
		{&Money{1234567, "EUR"}, "de_DE", "12.345,67"},
		{&Money{1234567890, "EUR"}, "de_DE", "12.345.678,90"},
		{&Money{-1, "EUR"}, "de_DE", "-0,01"},
		{&Money{-12, "EUR"}, "de_DE", "-0,12"},
		{&Money{-123, "EUR"}, "de_DE", "-1,23"},
		{&Money{-1234, "EUR"}, "de_DE", "-12,34"},
		{&Money{-123456, "EUR"}, "de_DE", "-1.234,56"},
		{&Money{-1234567, "EUR"}, "de_DE", "-12.345,67"},
		{&Money{-1234567890, "EUR"}, "de_DE", "-12.345.678,90"},

		{&Money{0, "EUR"}, "de_AT", "0,00"},
		{&Money{1, "EUR"}, "de_AT", "0,01"},
		{&Money{12, "EUR"}, "de_AT", "0,12"},
		{&Money{123, "EUR"}, "de_AT", "1,23"},
		{&Money{1234, "EUR"}, "de_AT", "12,34"},
		{&Money{123456, "EUR"}, "de_AT", "1.234,56"},
		{&Money{1234567, "EUR"}, "de_AT", "12.345,67"},
		{&Money{1234567890, "EUR"}, "de_AT", "12.345.678,90"},
		{&Money{-1, "EUR"}, "de_AT", "-0,01"},
		{&Money{-12, "EUR"}, "de_AT", "-0,12"},
		{&Money{-123, "EUR"}, "de_AT", "-1,23"},
		{&Money{-1234, "EUR"}, "de_AT", "-12,34"},
		{&Money{-123456, "EUR"}, "de_AT", "-1.234,56"},
		{&Money{-1234567, "EUR"}, "de_AT", "-12.345,67"},
		{&Money{-1234567890, "EUR"}, "de_AT", "-12.345.678,90"},

		{&Money{0, "EUR"}, "de_CH", "0.00"},
		{&Money{1, "EUR"}, "de_CH", "0.01"},
		{&Money{12, "EUR"}, "de_CH", "0.12"},
		{&Money{123, "EUR"}, "de_CH", "1.23"},
		{&Money{1234, "EUR"}, "de_CH", "12.34"},
		{&Money{123456, "EUR"}, "de_CH", "1'234.56"},
		{&Money{1234567, "EUR"}, "de_CH", "12'345.67"},
		{&Money{1234567890, "EUR"}, "de_CH", "12'345'678.90"},
		{&Money{-1, "EUR"}, "de_CH", "-0.01"},
		{&Money{-12, "EUR"}, "de_CH", "-0.12"},
		{&Money{-123, "EUR"}, "de_CH", "-1.23"},
		{&Money{-1234, "EUR"}, "de_CH", "-12.34"},
		{&Money{-123456, "EUR"}, "de_CH", "-1'234.56"},
		{&Money{-1234567, "EUR"}, "de_CH", "-12'345.67"},
		{&Money{-1234567890, "EUR"}, "de_CH", "-12'345'678.90"},

		{&Money{0, "EUR"}, "en", "0.00"},
		{&Money{1, "EUR"}, "en", "0.01"},
		{&Money{12, "EUR"}, "en", "0.12"},
		{&Money{123, "EUR"}, "en", "1.23"},
		{&Money{1234, "EUR"}, "en", "12.34"},
		{&Money{123456, "EUR"}, "en", "1234.56"},
		{&Money{1234567, "EUR"}, "en", "12345.67"},
		{&Money{1234567890, "EUR"}, "en", "12345678.90"},
		{&Money{-1, "EUR"}, "en", "-0.01"},
		{&Money{-12, "EUR"}, "en", "-0.12"},
		{&Money{-123, "EUR"}, "en", "-1.23"},
		{&Money{-1234, "EUR"}, "en", "-12.34"},
		{&Money{-123456, "EUR"}, "en", "-1234.56"},
		{&Money{-1234567, "EUR"}, "en", "-12345.67"},
		{&Money{-1234567890, "EUR"}, "en", "-12345678.90"},

		{&Money{0, "EUR"}, "en_US", "0.00"},
		{&Money{1, "EUR"}, "en_US", "0.01"},
		{&Money{12, "EUR"}, "en_US", "0.12"},
		{&Money{123, "EUR"}, "en_US", "1.23"},
		{&Money{1234, "EUR"}, "en_US", "12.34"},
		{&Money{123456, "EUR"}, "en_US", "1,234.56"},
		{&Money{1234567, "EUR"}, "en_US", "12,345.67"},
		{&Money{1234567890, "EUR"}, "en_US", "12,345,678.90"},
		{&Money{-1, "EUR"}, "en_US", "(0.01)"},
		{&Money{-12, "EUR"}, "en_US", "(0.12)"},
		{&Money{-123, "EUR"}, "en_US", "(1.23)"},
		{&Money{-1234, "EUR"}, "en_US", "(12.34)"},
		{&Money{-123456, "EUR"}, "en_US", "(1,234.56)"},
		{&Money{-1234567, "EUR"}, "en_US", "(12,345.67)"},
		{&Money{-1234567890, "EUR"}, "en_US", "(12,345,678.90)"},

		{&Money{0, "EUR"}, "fr", "0.00"},
		{&Money{1, "EUR"}, "fr", "0.01"},
		{&Money{12, "EUR"}, "fr", "0.12"},
		{&Money{123, "EUR"}, "fr", "1.23"},
		{&Money{1234, "EUR"}, "fr", "12.34"},
		{&Money{123456, "EUR"}, "fr", "1234.56"},
		{&Money{1234567, "EUR"}, "fr", "12345.67"},
		{&Money{1234567890, "EUR"}, "fr", "12345678.90"},
		{&Money{-1, "EUR"}, "fr", "-0.01"},
		{&Money{-12, "EUR"}, "fr", "-0.12"},
		{&Money{-123, "EUR"}, "fr", "-1.23"},
		{&Money{-1234, "EUR"}, "fr", "-12.34"},
		{&Money{-123456, "EUR"}, "fr", "-1234.56"},
		{&Money{-1234567, "EUR"}, "fr", "-12345.67"},
		{&Money{-1234567890, "EUR"}, "fr", "-12345678.90"},

		{&Money{1234567890, "USD"}, "en_US", "12,345,678.90"},
		{&Money{-1234567890, "USD"}, "en_US", "(12,345,678.90)"},
		{&Money{1234567890, "USD"}, "de_DE", "12.345.678,90"},
		{&Money{-1234567890, "USD"}, "de_DE", "-12.345.678,90"},
		{&Money{1234567890, "USD"}, "de_CH", "12'345'678.90"},
		{&Money{-1234567890, "USD"}, "de_CH", "-12'345'678.90"},
		{&Money{1234567890, "GBP"}, "en_GB", "12,345,678.90"},
		{&Money{-1234567890, "GBP"}, "en_GB", "-12,345,678.90"},
		{&Money{1234567890, "GBP"}, "de_DE", "12.345.678,90"},
		{&Money{-1234567890, "GBP"}, "de_DE", "-12.345.678,90"},
		{&Money{1234567890, "GBP"}, "de_CH", "12'345'678.90"},
		{&Money{-1234567890, "GBP"}, "de_CH", "-12'345'678.90"},
		{&Money{1234567890, "HUF"}, "hu_HU", "12 345 678,90"},
		{&Money{-1234567890, "HUF"}, "hu_HU", "-12 345 678,90"},
		{&Money{1234567890, "HUF"}, "de_DE", "12.345.678,90"},
		{&Money{-1234567890, "HUF"}, "de_DE", "-12.345.678,90"},
		{&Money{1234567890, "HUF"}, "de_CH", "12'345'678.90"},
		{&Money{-1234567890, "HUF"}, "de_CH", "-12'345'678.90"},
		{&Money{1234567890, "JPY"}, "ja_JP", "1,234,567,890"},
		{&Money{-1234567890, "JPY"}, "ja_JP", "-1,234,567,890"},
		{&Money{1234567890, "JPY"}, "de_DE", "1.234.567.890"},
		{&Money{-1234567890, "JPY"}, "de_DE", "-1.234.567.890"},
		{&Money{1234567890, "JPY"}, "de_CH", "1'234'567'890"},
		{&Money{-1234567890, "JPY"}, "de_CH", "-1'234'567'890"},
		{&Money{1234567890, "SEK"}, "se_SE", "12.345.678,90"},
		{&Money{-1234567890, "SEK"}, "se_SE", "-12.345.678,90"},
		{&Money{1234567890, "SEK"}, "de_DE", "12.345.678,90"},
		{&Money{-1234567890, "SEK"}, "de_DE", "-12.345.678,90"},
		{&Money{1234567890, "SEK"}, "de_CH", "12'345'678.90"},
		{&Money{-1234567890, "SEK"}, "de_CH", "-12'345'678.90"},

		{&Money{200000, "ZAR"}, "en_ZA", "2 000,00"},
		{&Money{120000, "ZAR"}, "en_ZA", "1 200,00"},
		{&Money{90000, "ZAR"}, "en_ZA", "900,00"},
		{&Money{90001, "ZAR"}, "en_ZA", "900,01"},
	}

	for _, f := range fixtures {
		got := f.m.FormatNoSymbol(f.locale)
		if got != f.expected {
			t.Errorf("expected '%s', got '%s' (currency: %s locale: %s)", f.expected, got, f.m.C, f.locale)
		}
	}
}

func TestMoneyFromMainUnit(t *testing.T) {
	var fixtures = []struct {
		m        string
		c        string
		expected *Money
	}{
		{"1", "GBP", &Money{100, "GBP"}},
		{"0.01", "GBP", &Money{1, "GBP"}},
		{"0.50", "GBP", &Money{50, "GBP"}},
		{"12.00", "GBP", &Money{1200, "GBP"}},
		{"120.99", "GBP", &Money{12099, "GBP"}},
		{"1.501", "GBP", &Money{150, "GBP"}},

		{"1.501", "BHD", &Money{1501, "BHD"}},
		{"1", "BND", &Money{1, "BND"}},

		{"1,500", "SEK", &Money{150000, "SEK"}},
	}
	for _, f := range fixtures {
		got := NewFromMainUnit(f.m, f.c)
		if !reflect.DeepEqual(got, f.expected) {
			t.Errorf("expected '%s', got '%s' (amount: %d currency: %s)", f.expected, got, got.M, got.C)
		}
	}
}
