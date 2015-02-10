// Copyright (c) 2011 Jad Dittmar
// See: https://github.com/Confunctionist/finance
//
// Some changes by Oliver Eilhard
package money

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/hailocab/i18n-go/currency"
	"github.com/hailocab/i18n-go/locale"
	"math"
	"strings"
)

type Money struct {
	M int64
	C string
}

var (
	ErrMoneyOverflow              = errors.New("i18n: money overflow")
	ErrMoneyDivideByZero          = errors.New("i18n: money division by zero")
	ErrMoneyDecimalPlacesTooLarge = errors.New("i18n: money decimal places too large")

	Guardi int     = 100
	Guard  int64   = int64(Guardi)
	Guardf float64 = float64(Guardi)
	DP     int64   = 100         // for default of 2 decimal places => 10^2 (can be reset)
	DPf    float64 = float64(DP) // for default of 2 decimal places => 10^2 (can be reset)
	Round          = .5
	Roundn         = Round * -1
)

const (
	MAXDEC = 18
)

func newDecimal(d int) int {
	if d < 0 {
		panic(ErrMoneyDivideByZero)
	}
	if d > MAXDEC {
		panic(ErrMoneyDecimalPlacesTooLarge)
	}
	var newDec int
	if d > 0 {
		newDec++
		for i := 0; i < d; i++ {
			newDec *= 10
		}
	}
	return newDec
}

// New returns a new Money that can be used for money arithmetic.
func New(m int64, c string) *Money {
	return &Money{m, c}
}

// Resets the package-wide decimal place (default is 2 decimal places).
func SetDecimal(d int) {
	decimal := newDecimal(d)
	DPf = float64(decimal)
	DP = int64(decimal)
	return
}

// Resets the package-wide decimal place by currency.
func SetDecimalByCurrency(cur string) {
	c := currency.Get(cur)
	if c != nil {
		decimal := newDecimal(c.DecimalDigits)
		DPf = float64(decimal)
		DP = int64(decimal)
	}
	return
}

// Resets the package-wide decimal place by locale.
func SetDecimalByLocale(lce string) {
	l := locale.Get(lce)
	if l != nil {
		decimal := newDecimal(l.CurrencyDecimalDigits)
		DPf = float64(decimal)
		DP = int64(decimal)
	}
	return
}

// Rounds int64 remainder rounded half towards plus infinity
// trunc = the remainder of the float64 calc
// r     = the result of the int64 cal
func Rnd(r int64, trunc float64) int64 {
	if trunc > 0 {
		if trunc >= Round {
			r++
		}
	} else {
		if trunc < Roundn {
			r--
		}
	}
	return r
}

// Returns the absolute value of Money.
func (m *Money) Abs() *Money {
	if m.M < 0 {
		m.Neg()
	}
	return m
}

// Adds two money types.
func (m *Money) Add(n *Money) *Money {
	r := m.M + n.M
	if (r^m.M)&(r^n.M) < 0 {
		panic(ErrMoneyOverflow)
	}
	m.M = r
	return m
}

// Divides one Money type from another.
func (m *Money) Div(n *Money) *Money {
	f := Guardf * DPf * float64(m.M) / float64(n.M) / Guardf
	i := int64(f)
	return m.Set(Rnd(i, f-float64(i)))
}

// Gets value of money truncating after DP (see Value() for no truncation).
func (m *Money) Gett() int64 {
	return m.M / DP
}

// Gets the float64 value of money (see Value() for int64).
func (m *Money) Get() float64 {
	return float64(m.M) / DPf
}

// Multiplies two Money types.
func (m *Money) Mul(n *Money) *Money {
	return m.Set(m.M * n.M / DP)
}

// Multiplies a Money with a float to return a money-stored type.
func (m *Money) Mulf(f float64) *Money {
	i := m.M * int64(f*Guardf*DPf)
	r := i / Guard / DP
	return m.Set(Rnd(r, float64(i)/Guardf/DPf-float64(r)))
}

// Returns the negative value of Money.
func (m *Money) Neg() *Money {
	if m.M != 0 {
		m.M *= -1
	}
	return m
}

// Sets the Money field M.
func (m *Money) Set(x int64) *Money {
	m.M = x
	return m
}

// Sets the currency of Money.
func (m *Money) SetCurrency(currency string) *Money {
	m.C = currency
	return m
}

// Sets the currency of Money by locale.
func (m *Money) SetCurrencyByLocale(lce string) *Money {
	l := locale.Get(lce)
	if l != nil {
		m.C = l.CurrencyCode
	}

	return m
}

// Sets the Money fields M and C.
func (m *Money) Setc(x int64, currency string) *Money {
	m.M = x
	m.C = currency
	return m
}

// Sets a float64 into a Money type for precision calculations.
func (m *Money) Setf(f float64) *Money {
	fDPf := f * DPf
	r := int64(f * DPf)
	return m.Set(Rnd(r, fDPf-float64(r)))
}

// Sets a float64 into a Money type for precision calculations.
func (m *Money) Setfc(f float64, currency string) *Money {
	fDPf := f * DPf
	r := int64(f * DPf)
	return m.Setc(Rnd(r, fDPf-float64(r)), currency)
}

// Returns the Sign of Money 1 if positive, -1 if negative.
func (m *Money) Sign() int {
	if m.M < 0 {
		return -1
	}
	return 1
}

// String for money type representation in basic monetary unit (DOLLARS CENTS).
func (m *Money) String() string {
	return m.toString(m.C)
}

// toString returns basic representation XX.XX CUR
func (m *Money) toString(curr string) string {
	if m.Sign() > 0 {
		return strings.TrimSpace(fmt.Sprintf("%d.%02d %s", m.Value()/DP, m.Value()%DP, curr))
	}
	// Negative value
	return strings.TrimSpace(fmt.Sprintf("-%d.%02d %s", m.Abs().Value()/DP, m.Abs().Value()%DP, curr))
}

func (m *Money) Format(loc string) string {
	// DP is a measure for decimals: 2 decimal digits => dp = 10^2
	currencySymbol := m.C
	if curr := currency.Get(m.C); curr != nil {
		currencySymbol = curr.Symbol
	}

	return m.format(loc, currencySymbol)
}

func (m *Money) format(loc, symbol string) string {
	l := locale.Get(loc)
	if l == nil {
		// If we don't have any information about the currency format,
		// we'll try our best to display something useful.
		if len(symbol) == 0 {
			return m.toString("") // force to no symbol
		}
		return m.String()
	}

	// DP is a measure for decimals: 2 decimal digits => dp = 10^2
	dp := int64(math.Pow10(l.CurrencyDecimalDigits))

	// Group DP is a measure for grouping: 3 decimal digits => groupDp = 10^3
	groupSize := 3
	if len(l.CurrencyGroupSizes) >= 1 {
		// BUG(oe): Handle currency group size
		groupSize = l.CurrencyGroupSizes[0]
	}
	groupDp := int64(math.Pow10(groupSize))

	// We use absolute values (as int64) from here on, because the
	// negative sign is part of the currency format pattern.
	absVal := m.Value()
	if m.Sign() < 0 {
		absVal = -absVal
	}
	wholeVal := absVal / dp
	decVal := absVal % dp

	// The unformatted string (without grouping and with a decimal sep of ".")
	var unformatted string
	if l.CurrencyDecimalDigits > 0 {
		unformatted = fmt.Sprintf("%d.%0"+fmt.Sprintf("%d", l.CurrencyDecimalDigits)+"d", wholeVal, decVal)
	} else {
		unformatted = fmt.Sprintf("%d", wholeVal)
	}

	// Perform grouping operation of the whole number
	groups := make([]string, 0)
	innerGroupFmt := "%0" + fmt.Sprintf("%d", groupSize) + "d"
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
			wholeBuf.WriteString(l.CurrencyGroupSeparator)
		}
		wholeBuf.WriteString(groups[len(groups)-i-1])
	}

	// Which pattern do we need?
	// Notice that the minus sign is part of the pattern
	var pattern string
	if m.Sign() > 0 {
		pattern = l.CurrencyPositivePattern
	} else {
		pattern = l.CurrencyNegativePattern
	}

	// Split into whole and decimal and build formatted number
	var formatted string
	parts := strings.SplitN(unformatted, ".", 2)
	if len(parts) > 1 {
		formatted = fmt.Sprintf("%s%s%s", wholeBuf.String(), l.CurrencyDecimalSeparator, parts[1])
	} else {
		formatted = wholeBuf.String()
	}
	output := strings.Replace(pattern, "$", symbol, -1)
	output = strings.Replace(output, "n", formatted, -1)

	return output
}

func (m *Money) FormatNoSymbol(loc string) string {
	ret := strings.TrimSpace(m.format(loc, ""))
	return strings.Replace(ret, "- ", "-", -1) // remove any funny whitespace due to -ve pattern
}

// Subtracts one Money type from another.
func (m *Money) Sub(n *Money) *Money {
	r := m.M - n.M
	if (r^m.M)&^(r^n.M) < 0 {
		panic(ErrMoneyOverflow)
	}
	m.M = r
	return m
}

// Returns in int64 the value of Money (also see Gett(), See Get() for float64).
func (m *Money) Value() int64 {
	return m.M
}
