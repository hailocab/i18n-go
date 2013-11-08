package territory

import (
	"testing"
)

func TestTerritories(t *testing.T) {
	var tests = []struct {
		code                string
		found               bool
		expectedEnglishName string
		expectedNativeName  string
	}{
		/* 0 */ {"XY", false, "", ""},
		/* 1 */ {"DE", true, "Germany", "Deutschland"},
		// BUG(oe): Wrong native name, at least from a German standpoint ;-)
		/* 2 */ {"CH", true, "Switzerland", "Svizra"},
		// BUG(oe): Wrong native name, at least from a German standpoint ;-)
		/* 3 */ {"GB", true, "United Kingdom", "y Deyrnas Unedig"},
		/* 4 */ {"US", true, "United States", "United States"},
	}

	for i, f := range tests {
		c, found := territories[f.code]
		if found != f.found {
			t.Fatalf("%d. expected territory %s found flag to be %v, got %v", i, f.code, f.found, found)
		}
		if f.found {
			if c == nil {
				t.Fatalf("%d. expected country to be != nil", i)
			}
			if f.expectedEnglishName != c.EnglishName {
				t.Errorf("%d. expected EnglishName to be %v, got %v", i, f.expectedEnglishName, c.EnglishName)
			}
			if f.expectedNativeName != c.NativeName {
				t.Errorf("%d. expected NativeName to be %v, got %v", i, f.expectedNativeName, c.NativeName)
			}
		}
	}
}
