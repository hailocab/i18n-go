package language

// Language represents all information about a language.
type Language struct {
	// Code is the 2-letter, downcase ISO code of the language (e.g. de).
	Code string
	// NativeName is the name of the language in the language itself.
	NativeName string
	// EnglishName is the name of the language in English.
	EnglishName string
}

func Get(code string) *Language {
	return languages[code]
}

func Languages() map[string]*Language {
	return languages
}
