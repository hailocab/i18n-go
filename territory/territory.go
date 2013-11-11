package territory

// Territory represents information about a country.
type Territory struct {
	// Code is the upcase ISO-code of the territory.
	Code string
	// NativeName is the name of the territory in the specified territory.
	NativeName string
	// EnglishName is the English name of the territory.
	EnglishName string
}

func Get(code string) *Territory {
	return territories[code]
}

func Territories() map[string]*Territory {
	return territories
}
