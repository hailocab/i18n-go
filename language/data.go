package language

var languages = map[string]*Language{
	"af": &Language{
		Code:        "af",
		NativeName:  "Afrikaans",
		EnglishName: "Afrikaans",
	},
	"am": &Language{
		Code:        "am",
		NativeName:  "አማርኛ",
		EnglishName: "Amharic",
	},
	"ar": &Language{
		Code:        "ar",
		NativeName:  "العربية",
		EnglishName: "Arabic",
	},
	"arn": &Language{
		Code:        "arn",
		NativeName:  "Mapudungun",
		EnglishName: "Mapudungun",
	},
	"as": &Language{
		Code:        "as",
		NativeName:  "অসমীয়া",
		EnglishName: "Assamese",
	},
	"az": &Language{
		Code:        "az",
		NativeName:  "Azərbaycan­ılı",
		EnglishName: "Azeri",
	},
	"ba": &Language{
		Code:        "ba",
		NativeName:  "Башҡорт",
		EnglishName: "Bashkir",
	},
	"be": &Language{
		Code:        "be",
		NativeName:  "Беларускі",
		EnglishName: "Belarusian",
	},
	"bg": &Language{
		Code:        "bg",
		NativeName:  "български",
		EnglishName: "Bulgarian",
	},
	"bn": &Language{
		Code:        "bn",
		NativeName:  "বাংলা",
		EnglishName: "Bengali",
	},
	"bo": &Language{
		Code:        "bo",
		NativeName:  "བོད་ཡིག",
		EnglishName: "Tibetan",
	},
	"br": &Language{
		Code:        "br",
		NativeName:  "brezhoneg",
		EnglishName: "Breton",
	},
	"bs": &Language{
		Code:        "bs",
		NativeName:  "босански",
		EnglishName: "Bosnian (Cyrillic)",
	},
	"ca": &Language{
		Code:        "ca",
		NativeName:  "català",
		EnglishName: "Catalan",
	},
	"co": &Language{
		Code:        "co",
		NativeName:  "Corsu",
		EnglishName: "Corsican",
	},
	"cs": &Language{
		Code:        "cs",
		NativeName:  "čeština",
		EnglishName: "Czech",
	},
	"cy": &Language{
		Code:        "cy",
		NativeName:  "Cymraeg",
		EnglishName: "Welsh",
	},
	"da": &Language{
		Code:        "da",
		NativeName:  "dansk",
		EnglishName: "Danish",
	},
	"de": &Language{
		Code:        "de",
		NativeName:  "Deutsch",
		EnglishName: "German",
	},
	"dsb": &Language{
		Code:        "dsb",
		NativeName:  "dolnoserbšćina",
		EnglishName: "Lower Sorbian",
	},
	"dv": &Language{
		Code:        "dv",
		NativeName:  "ދިވެހިބަސް",
		EnglishName: "Divehi",
	},
	"el": &Language{
		Code:        "el",
		NativeName:  "Ελληνικά",
		EnglishName: "Greek",
	},
	"en": &Language{
		Code:        "en",
		NativeName:  "English",
		EnglishName: "English",
	},
	"es": &Language{
		Code:        "es",
		NativeName:  "español",
		EnglishName: "Spanish",
	},
	"et": &Language{
		Code:        "et",
		NativeName:  "eesti",
		EnglishName: "Estonian",
	},
	"eu": &Language{
		Code:        "eu",
		NativeName:  "euskara",
		EnglishName: "Basque",
	},
	"fa": &Language{
		Code:        "fa",
		NativeName:  "فارسى",
		EnglishName: "Persian",
	},
	"fi": &Language{
		Code:        "fi",
		NativeName:  "suomi",
		EnglishName: "Finnish",
	},
	"fil": &Language{
		Code:        "fil",
		NativeName:  "Filipino",
		EnglishName: "Filipino",
	},
	"fo": &Language{
		Code:        "fo",
		NativeName:  "føroyskt",
		EnglishName: "Faroese",
	},
	"fr": &Language{
		Code:        "fr",
		NativeName:  "français",
		EnglishName: "French",
	},
	"fy": &Language{
		Code:        "fy",
		NativeName:  "Frysk",
		EnglishName: "Frisian",
	},
	"ga": &Language{
		Code:        "ga",
		NativeName:  "Gaeilge",
		EnglishName: "Irish",
	},
	"gd": &Language{
		Code:        "gd",
		NativeName:  "Gàidhlig",
		EnglishName: "Scottish Gaelic",
	},
	"gl": &Language{
		Code:        "gl",
		NativeName:  "galego",
		EnglishName: "Galician",
	},
	"gsw": &Language{
		Code:        "gsw",
		NativeName:  "Elsässisch",
		EnglishName: "Alsatian",
	},
	"gu": &Language{
		Code:        "gu",
		NativeName:  "ગુજરાતી",
		EnglishName: "Gujarati",
	},
	"ha": &Language{
		Code:        "ha",
		NativeName:  "Hausa",
		EnglishName: "Hausa",
	},
	"he": &Language{
		Code:        "he",
		NativeName:  "עברית",
		EnglishName: "Hebrew",
	},
	"hi": &Language{
		Code:        "hi",
		NativeName:  "हिंदी",
		EnglishName: "Hindi",
	},
	"hr": &Language{
		Code:        "hr",
		NativeName:  "hrvatski",
		EnglishName: "Croatian",
	},
	"hsb": &Language{
		Code:        "hsb",
		NativeName:  "hornjoserbšćina",
		EnglishName: "Upper Sorbian",
	},
	"hu": &Language{
		Code:        "hu",
		NativeName:  "magyar",
		EnglishName: "Hungarian",
	},
	"hy": &Language{
		Code:        "hy",
		NativeName:  "Հայերեն",
		EnglishName: "Armenian",
	},
	"id": &Language{
		Code:        "id",
		NativeName:  "Bahasa Indonesia",
		EnglishName: "Indonesian",
	},
	"ig": &Language{
		Code:        "ig",
		NativeName:  "Igbo",
		EnglishName: "Igbo",
	},
	"ii": &Language{
		Code:        "ii",
		NativeName:  "ꆈꌠꁱꂷ",
		EnglishName: "Yi",
	},
	"is": &Language{
		Code:        "is",
		NativeName:  "íslenska",
		EnglishName: "Icelandic",
	},
	"it": &Language{
		Code:        "it",
		NativeName:  "italiano",
		EnglishName: "Italian",
	},
	"iu": &Language{
		Code:        "iu",
		NativeName:  "Inuktitut",
		EnglishName: "Inuktitut",
	},
	"iv": &Language{
		Code:        "iv",
		NativeName:  "Invariant Language (Invariant Country)",
		EnglishName: "Invariant Language (Invariant Country)",
	},
	"ja": &Language{
		Code:        "ja",
		NativeName:  "日本語",
		EnglishName: "Japanese",
	},
	"ka": &Language{
		Code:        "ka",
		NativeName:  "ქართული",
		EnglishName: "Georgian",
	},
	"kk": &Language{
		Code:        "kk",
		NativeName:  "Қазақ",
		EnglishName: "Kazakh",
	},
	"kl": &Language{
		Code:        "kl",
		NativeName:  "kalaallisut",
		EnglishName: "Greenlandic",
	},
	"km": &Language{
		Code:        "km",
		NativeName:  "ខ្មែរ",
		EnglishName: "Khmer",
	},
	"kn": &Language{
		Code:        "kn",
		NativeName:  "ಕನ್ನಡ",
		EnglishName: "Kannada",
	},
	"ko": &Language{
		Code:        "ko",
		NativeName:  "한국어",
		EnglishName: "Korean",
	},
	"kok": &Language{
		Code:        "kok",
		NativeName:  "कोंकणी",
		EnglishName: "Konkani",
	},
	"ky": &Language{
		Code:        "ky",
		NativeName:  "Кыргыз",
		EnglishName: "Kyrgyz",
	},
	"lb": &Language{
		Code:        "lb",
		NativeName:  "Lëtzebuergesch",
		EnglishName: "Luxembourgish",
	},
	"lo": &Language{
		Code:        "lo",
		NativeName:  "ລາວ",
		EnglishName: "Lao",
	},
	"lt": &Language{
		Code:        "lt",
		NativeName:  "lietuvių",
		EnglishName: "Lithuanian",
	},
	"lv": &Language{
		Code:        "lv",
		NativeName:  "latviešu",
		EnglishName: "Latvian",
	},
	"mi": &Language{
		Code:        "mi",
		NativeName:  "Reo Māori",
		EnglishName: "Maori",
	},
	"mk": &Language{
		Code:        "mk",
		NativeName:  "македонски јазик",
		EnglishName: "Macedonian (FYROM)",
	},
	"ml": &Language{
		Code:        "ml",
		NativeName:  "മലയാളം",
		EnglishName: "Malayalam",
	},
	"mn": &Language{
		Code:        "mn",
		NativeName:  "Монгол хэл",
		EnglishName: "Mongolian",
	},
	"moh": &Language{
		Code:        "moh",
		NativeName:  "Kanien'kéha",
		EnglishName: "Mohawk",
	},
	"mr": &Language{
		Code:        "mr",
		NativeName:  "मराठी",
		EnglishName: "Marathi",
	},
	"ms": &Language{
		Code:        "ms",
		NativeName:  "Bahasa Melayu",
		EnglishName: "Malay",
	},
	"mt": &Language{
		Code:        "mt",
		NativeName:  "Malti",
		EnglishName: "Maltese",
	},
	"nb": &Language{
		Code:        "nb",
		NativeName:  "norsk",
		EnglishName: "Norwegian",
	},
	"ne": &Language{
		Code:        "ne",
		NativeName:  "नेपाली",
		EnglishName: "Nepali",
	},
	"nl": &Language{
		Code:        "nl",
		NativeName:  "Nederlands",
		EnglishName: "Dutch",
	},
	"nn": &Language{
		Code:        "nn",
		NativeName:  "norsk (nynorsk)",
		EnglishName: "Norwegian (Nynorsk)",
	},
	"nso": &Language{
		Code:        "nso",
		NativeName:  "Sesotho sa Leboa",
		EnglishName: "Sesotho sa Leboa",
	},
	"oc": &Language{
		Code:        "oc",
		NativeName:  "Occitan",
		EnglishName: "Occitan",
	},
	"or": &Language{
		Code:        "or",
		NativeName:  "ଓଡ଼ିଆ",
		EnglishName: "Oriya",
	},
	"pa": &Language{
		Code:        "pa",
		NativeName:  "ਪੰਜਾਬੀ",
		EnglishName: "Punjabi",
	},
	"pl": &Language{
		Code:        "pl",
		NativeName:  "polski",
		EnglishName: "Polish",
	},
	"prs": &Language{
		Code:        "prs",
		NativeName:  "درى",
		EnglishName: "Dari",
	},
	"ps": &Language{
		Code:        "ps",
		NativeName:  "پښتو",
		EnglishName: "Pashto",
	},
	"pt": &Language{
		Code:        "pt",
		NativeName:  "Português",
		EnglishName: "Portuguese",
	},
	"qut": &Language{
		Code:        "qut",
		NativeName:  "K'iche",
		EnglishName: "K'iche",
	},
	"quz": &Language{
		Code:        "quz",
		NativeName:  "runasimi",
		EnglishName: "Quechua",
	},
	"rm": &Language{
		Code:        "rm",
		NativeName:  "Rumantsch",
		EnglishName: "Romansh",
	},
	"ro": &Language{
		Code:        "ro",
		NativeName:  "română",
		EnglishName: "Romanian",
	},
	"ru": &Language{
		Code:        "ru",
		NativeName:  "русский",
		EnglishName: "Russian",
	},
	"rw": &Language{
		Code:        "rw",
		NativeName:  "Kinyarwanda",
		EnglishName: "Kinyarwanda",
	},
	"sa": &Language{
		Code:        "sa",
		NativeName:  "संस्कृत",
		EnglishName: "Sanskrit",
	},
	"sah": &Language{
		Code:        "sah",
		NativeName:  "саха",
		EnglishName: "Yakut",
	},
	"se": &Language{
		Code:        "se",
		NativeName:  "davvisámegiella",
		EnglishName: "Sami (Northern)",
	},
	"si": &Language{
		Code:        "si",
		NativeName:  "සිංහල",
		EnglishName: "Sinhala",
	},
	"sk": &Language{
		Code:        "sk",
		NativeName:  "slovenčina",
		EnglishName: "Slovak",
	},
	"sl": &Language{
		Code:        "sl",
		NativeName:  "slovenski",
		EnglishName: "Slovenian",
	},
	"sma": &Language{
		Code:        "sma",
		NativeName:  "åarjelsaemiengiele",
		EnglishName: "Sami (Southern)",
	},
	"smj": &Language{
		Code:        "smj",
		NativeName:  "julevusámegiella",
		EnglishName: "Sami (Lule)",
	},
	"smn": &Language{
		Code:        "smn",
		NativeName:  "sämikielâ",
		EnglishName: "Sami (Inari)",
	},
	"sms": &Language{
		Code:        "sms",
		NativeName:  "sääm´ǩiõll",
		EnglishName: "Sami (Skolt)",
	},
	"sq": &Language{
		Code:        "sq",
		NativeName:  "shqipe",
		EnglishName: "Albanian",
	},
	"sr": &Language{
		Code:        "sr",
		NativeName:  "српски",
		EnglishName: "Serbian (Cyrillic)",
	},
	"sv": &Language{
		Code:        "sv",
		NativeName:  "svenska",
		EnglishName: "Swedish",
	},
	"sw": &Language{
		Code:        "sw",
		NativeName:  "Kiswahili",
		EnglishName: "Kiswahili",
	},
	"syr": &Language{
		Code:        "syr",
		NativeName:  "ܣܘܪܝܝܐ",
		EnglishName: "Syriac",
	},
	"ta": &Language{
		Code:        "ta",
		NativeName:  "தமிழ்",
		EnglishName: "Tamil",
	},
	"te": &Language{
		Code:        "te",
		NativeName:  "తెలుగు",
		EnglishName: "Telugu",
	},
	"tg": &Language{
		Code:        "tg",
		NativeName:  "Тоҷикӣ",
		EnglishName: "Tajik",
	},
	"th": &Language{
		Code:        "th",
		NativeName:  "ไทย",
		EnglishName: "Thai",
	},
	"tk": &Language{
		Code:        "tk",
		NativeName:  "türkmençe",
		EnglishName: "Turkmen",
	},
	"tn": &Language{
		Code:        "tn",
		NativeName:  "Setswana",
		EnglishName: "Setswana",
	},
	"tr": &Language{
		Code:        "tr",
		NativeName:  "Türkçe",
		EnglishName: "Turkish",
	},
	"tt": &Language{
		Code:        "tt",
		NativeName:  "Татар",
		EnglishName: "Tatar",
	},
	"tzm": &Language{
		Code:        "tzm",
		NativeName:  "Tamazight",
		EnglishName: "Tamazight",
	},
	"ug": &Language{
		Code:        "ug",
		NativeName:  "ئۇيغۇرچە",
		EnglishName: "Uyghur",
	},
	"uk": &Language{
		Code:        "uk",
		NativeName:  "українська",
		EnglishName: "Ukrainian",
	},
	"ur": &Language{
		Code:        "ur",
		NativeName:  "اُردو",
		EnglishName: "Urdu",
	},
	"uz": &Language{
		Code:        "uz",
		NativeName:  "U'zbek",
		EnglishName: "Uzbek",
	},
	"vi": &Language{
		Code:        "vi",
		NativeName:  "Tiếng Việt",
		EnglishName: "Vietnamese",
	},
	"wo": &Language{
		Code:        "wo",
		NativeName:  "Wolof",
		EnglishName: "Wolof",
	},
	"xh": &Language{
		Code:        "xh",
		NativeName:  "isiXhosa",
		EnglishName: "isiXhosa",
	},
	"yo": &Language{
		Code:        "yo",
		NativeName:  "Yoruba",
		EnglishName: "Yoruba",
	},
	"zh": &Language{
		Code:        "zh",
		NativeName:  "中文(简体)",
		EnglishName: "Chinese (Simplified)",
	},
	"zu": &Language{
		Code:        "zu",
		NativeName:  "isiZulu",
		EnglishName: "isiZulu",
	},
}
