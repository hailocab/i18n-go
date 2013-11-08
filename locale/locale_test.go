package locale

import (
	"testing"
)

func TestLocales(t *testing.T) {
	var tests = []struct {
		code              string
		found             bool
		expectedLanguage  string
		expectedLangISO3  string
		expectedTerritory string
		expectedCurrCode  string
	}{
		/*  0 */ {"xy", false, "", "", "", ""},
		/*  1 */ {"de", false, "", "", "", ""},
		/*  2 */ {"en", false, "", "", "", ""},
		/*  3 */ {"fr", false, "", "", "", ""},
		/*  4 */ {"de_AT", true, "de", "deu", "AT", "EUR"},
		/*  5 */ {"de_CH", true, "de", "deu", "CH", "CHF"},
		/*  6 */ {"de_DE", true, "de", "deu", "DE", "EUR"},
		/*  7 */ {"en_CA", true, "en", "eng", "CA", "CAD"},
		/*  8 */ {"en_GB", true, "en", "eng", "GB", "GBP"},
		/*  9 */ {"en_US", true, "en", "eng", "US", "USD"},
		/* 10 */ {"fr_CH", true, "fr", "fra", "CH", "CHF"},
		/* 11 */ {"fr_FR", true, "fr", "fra", "FR", "EUR"},
	}

	for i, f := range tests {
		loc, found := locales[f.code]
		if found != f.found {
			t.Fatalf("%d. expected %v found flag to be %v, got %v", i, f.code, f.found, found)
		}
		if found {
			if loc == nil {
				t.Fatalf("%d. expected locale to be != nil", i)
			}
			if f.expectedLanguage != loc.Language {
				t.Errorf("%d. expected Language to be %v, got %v", i, f.expectedLanguage, loc.Language)
			}
			if f.expectedLangISO3 != loc.LanguageISO3 {
				t.Errorf("%d. expected LanguageISO3 to be %v, got %v", i, f.expectedLangISO3, loc.LanguageISO3)
			}
			if f.expectedTerritory != loc.Territory {
				t.Errorf("%d. expected Territory to be %v, got %v", i, f.expectedTerritory, loc.Territory)
			}
			if f.expectedCurrCode != loc.CurrencyCode {
				t.Errorf("%d. expected CurrencyCode to be %v, got %v", i, f.expectedCurrCode, loc.CurrencyCode)
			}
		}
	}
}
