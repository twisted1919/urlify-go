package urlify

import (
	"regexp"
	"strings"
)

var charsMap = map[string]map[string]string{
	"de": map[string]string{ /* German */
		"Ä": "Ae", "Ö": "Oe", "Ü": "Ue", "ä": "ae", "ö": "oe", "ü": "ue", "ß": "ss",
		"ẞ": "SS",
	},
	"latin": map[string]string{
		"À": "A", "Á": "A", "Â": "A", "Ã": "A", "Ä": "A", "Å": "A", "Ă": "A", "Æ": "AE", "Ç": "C", "È": "E",
		"É": "E", "Ê": "E", "Ë": "E", "Ì": "I", "Í": "I", "Î": "I",
		"Ï": "I", "Ð": "D", "Ñ": "N", "Ò": "O", "Ó": "O", "Ô": "O", "Õ": "O", "Ö": "O", "Ő": "O", "Ø": "O",
		"Ș": "S", "Ț": "T", "Ù": "U", "Ú": "U", "Û": "U", "Ü": "U", "Ű": "U",
		"Ý": "Y", "Þ": "TH", "ß": "ss", "à": "a", "á": "a", "â": "a", "ã": "a", "ä": "a", "å": "a", "ă": "a",
		"æ": "ae", "ç": "c", "è": "e", "é": "e", "ê": "e", "ë": "e",
		"ì": "i", "í": "i", "î": "i", "ï": "i", "ð": "d", "ñ": "n", "ò": "o", "ó": "o", "ô": "o", "õ": "o",
		"ö": "o", "ő": "o", "ø": "o", "ș": "s", "ț": "t", "ù": "u", "ú": "u",
		"û": "u", "ü": "u", "ű": "u", "ý": "y", "þ": "th", "ÿ": "y",
	},
	"latin_symbols": map[string]string{
		"©": "(c)",
	},
	"el": map[string]string{ /* Greek */
		"α": "a", "β": "b", "γ": "g", "δ": "d", "ε": "e", "ζ": "z", "η": "h", "θ": "8",
		"ι": "i", "κ": "k", "λ": "l", "μ": "m", "ν": "n", "ξ": "3", "ο": "o", "π": "p",
		"ρ": "r", "σ": "s", "τ": "t", "υ": "y", "φ": "f", "χ": "x", "ψ": "ps", "ω": "w",
		"ά": "a", "έ": "e", "ί": "i", "ό": "o", "ύ": "y", "ή": "h", "ώ": "w", "ς": "s",
		"ϊ": "i", "ΰ": "y", "ϋ": "y", "ΐ": "i",
		"Α": "A", "Β": "B", "Γ": "G", "Δ": "D", "Ε": "E", "Ζ": "Z", "Η": "H", "Θ": "8",
		"Ι": "I", "Κ": "K", "Λ": "L", "Μ": "M", "Ν": "N", "Ξ": "3", "Ο": "O", "Π": "P",
		"Ρ": "R", "Σ": "S", "Τ": "T", "Υ": "Y", "Φ": "F", "Χ": "X", "Ψ": "PS", "Ω": "W",
		"Ά": "A", "Έ": "E", "Ί": "I", "Ό": "O", "Ύ": "Y", "Ή": "H", "Ώ": "W", "Ϊ": "I",
		"Ϋ": "Y",
	},
	"tr": map[string]string{ /* Turkish */
		"ş": "s", "Ş": "S", "ı": "i", "İ": "I", "ç": "c", "Ç": "C", "ü": "u", "Ü": "U",
		"ö": "o", "Ö": "O", "ğ": "g", "Ğ": "G",
	},
	"bg": map[string]string{ /* Bulgarian */
		"Щ": "Sht", "Ш": "Sh", "Ч": "Ch", "Ц": "C", "Ю": "Yu", "Я": "Ya",
		"Ж": "J", "А": "A", "Б": "B", "В": "V", "Г": "G", "Д": "D",
		"Е": "E", "З": "Z", "И": "I", "Й": "Y", "К": "K", "Л": "L",
		"М": "M", "Н": "N", "О": "O", "П": "P", "Р": "R", "С": "S",
		"Т": "T", "У": "U", "Ф": "F", "Х": "H", "Ь": "", "Ъ": "A",
		"щ": "sht", "ш": "sh", "ч": "ch", "ц": "c", "ю": "yu", "я": "ya",
		"ж": "j", "а": "a", "б": "b", "в": "v", "г": "g", "д": "d",
		"е": "e", "з": "z", "и": "i", "й": "y", "к": "k", "л": "l",
		"м": "m", "н": "n", "о": "o", "п": "p", "р": "r", "с": "s",
		"т": "t", "у": "u", "ф": "f", "х": "h", "ь": "", "ъ": "a",
	},
	"ru": map[string]string{ /* Russian */
		"а": "a", "б": "b", "в": "v", "г": "g", "д": "d", "е": "e", "ё": "yo", "ж": "zh",
		"з": "z", "и": "i", "й": "j", "к": "k", "л": "l", "м": "m", "н": "n", "о": "o",
		"п": "p", "р": "r", "с": "s", "т": "t", "у": "u", "ф": "f", "х": "h", "ц": "c",
		"ч": "ch", "ш": "sh", "щ": "sh", "ъ": "", "ы": "y", "ь": "", "э": "e", "ю": "yu",
		"я": "ya",
		"А": "A", "Б": "B", "В": "V", "Г": "G", "Д": "D", "Е": "E", "Ё": "Yo", "Ж": "Zh",
		"З": "Z", "И": "I", "Й": "J", "К": "K", "Л": "L", "М": "M", "Н": "N", "О": "O",
		"П": "P", "Р": "R", "С": "S", "Т": "T", "У": "U", "Ф": "F", "Х": "H", "Ц": "C",
		"Ч": "Ch", "Ш": "Sh", "Щ": "Sh", "Ъ": "", "Ы": "Y", "Ь": "", "Э": "E", "Ю": "Yu",
		"Я": "Ya",
		"№": "",
	},
	"uk": map[string]string{ /* Ukrainian */
		"Є": "Ye", "І": "I", "Ї": "Yi", "Ґ": "G", "є": "ye", "і": "i", "ї": "yi", "ґ": "g",
	},
	"cs": map[string]string{ /* Czech */
		"č": "c", "ď": "d", "ě": "e", "ň": "n", "ř": "r", "š": "s", "ť": "t", "ů": "u",
		"ž": "z", "Č": "C", "Ď": "D", "Ě": "E", "Ň": "N", "Ř": "R", "Š": "S", "Ť": "T",
		"Ů": "U", "Ž": "Z",
	},
	"pl": map[string]string{ /* Polish */
		"ą": "a", "ć": "c", "ę": "e", "ł": "l", "ń": "n", "ó": "o", "ś": "s", "ź": "z",
		"ż": "z", "Ą": "A", "Ć": "C", "Ę": "e", "Ł": "L", "Ń": "N", "Ó": "O", "Ś": "S",
		"Ź": "Z", "Ż": "Z",
	},
	"ro": map[string]string{ /* Romanian */
		"ă": "a", "â": "a", "î": "i", "ș": "s", "ț": "t", "Ţ": "T", "ţ": "t",
	},
	"lv": map[string]string{ /* Latvian */
		"ā": "a", "č": "c", "ē": "e", "ģ": "g", "ī": "i", "ķ": "k", "ļ": "l", "ņ": "n",
		"š": "s", "ū": "u", "ž": "z", "Ā": "A", "Č": "C", "Ē": "E", "Ģ": "G", "Ī": "i",
		"Ķ": "k", "Ļ": "L", "Ņ": "N", "Š": "S", "Ū": "u", "Ž": "Z",
	},
	"lt": map[string]string{ /* Lithuanian */
		"ą": "a", "č": "c", "ę": "e", "ė": "e", "į": "i", "š": "s", "ų": "u", "ū": "u", "ž": "z",
		"Ą": "A", "Č": "C", "Ę": "E", "Ė": "E", "Į": "I", "Š": "S", "Ų": "U", "Ū": "U", "Ž": "Z",
	},
	"vn": map[string]string{ /* Vietnamese */
		"Á": "A", "À": "A", "Ả": "A", "Ã": "A", "Ạ": "A", "Ă": "A", "Ắ": "A", "Ằ": "A", "Ẳ": "A", "Ẵ": "A",
		"Ặ": "A", "Â": "A", "Ấ": "A", "Ầ": "A", "Ẩ": "A", "Ẫ": "A", "Ậ": "A",
		"á": "a", "à": "a", "ả": "a", "ã": "a", "ạ": "a", "ă": "a", "ắ": "a", "ằ": "a", "ẳ": "a", "ẵ": "a",
		"ặ": "a", "â": "a", "ấ": "a", "ầ": "a", "ẩ": "a", "ẫ": "a", "ậ": "a",
		"É": "E", "È": "E", "Ẻ": "E", "Ẽ": "E", "Ẹ": "E", "Ê": "E", "Ế": "E", "Ề": "E", "Ể": "E", "Ễ": "E",
		"Ệ": "E",
		"é": "e", "è": "e", "ẻ": "e", "ẽ": "e", "ẹ": "e", "ê": "e", "ế": "e", "ề": "e", "ể": "e", "ễ": "e",
		"ệ": "e",
		"Í": "I", "Ì": "I", "Ỉ": "I", "Ĩ": "I", "Ị": "I", "í": "i", "ì": "i", "ỉ": "i", "ĩ": "i", "ị": "i",
		"Ó": "O", "Ò": "O", "Ỏ": "O", "Õ": "O", "Ọ": "O", "Ô": "O", "Ố": "O", "Ồ": "O", "Ổ": "O", "Ỗ": "O",
		"Ộ": "O", "Ơ": "O", "Ớ": "O", "Ờ": "O", "Ở": "O", "Ỡ": "O", "Ợ": "O",
		"ó": "o", "ò": "o", "ỏ": "o", "õ": "o", "ọ": "o", "ô": "o", "ố": "o", "ồ": "o", "ổ": "o", "ỗ": "o",
		"ộ": "o", "ơ": "o", "ớ": "o", "ờ": "o", "ở": "o", "ỡ": "o", "ợ": "o",
		"Ú": "U", "Ù": "U", "Ủ": "U", "Ũ": "U", "Ụ": "U", "Ư": "U", "Ứ": "U", "Ừ": "U", "Ử": "U", "Ữ": "U",
		"Ự": "U",
		"ú": "u", "ù": "u", "ủ": "u", "ũ": "u", "ụ": "u", "ư": "u", "ứ": "u", "ừ": "u", "ử": "u", "ữ": "u",
		"ự": "u",
		"Ý": "Y", "Ỳ": "Y", "Ỷ": "Y", "Ỹ": "Y", "Ỵ": "Y", "ý": "y", "ỳ": "y", "ỷ": "y", "ỹ": "y", "ỵ": "y",
		"Đ": "D", "đ": "d",
	},
	"ar": map[string]string{ /* Arabic */
		"أ": "a", "ب": "b", "ت": "t", "ث": "th", "ج": "g", "ح": "h", "خ": "kh", "د": "d",
		"ذ": "th", "ر": "r", "ز": "z", "س": "s", "ش": "sh", "ص": "s", "ض": "d", "ط": "t",
		"ظ": "th", "ع": "aa", "غ": "gh", "ف": "f", "ق": "k", "ك": "k", "ل": "l", "م": "m",
		"ن": "n", "ه": "h", "و": "o", "ي": "y",
	},
	"sr": map[string]string{ /* Serbian */
		"ђ": "dj", "ј": "j", "љ": "lj", "њ": "nj", "ћ": "c", "џ": "dz", "đ": "dj",
		"Ђ": "Dj", "Ј": "j", "Љ": "Lj", "Њ": "Nj", "Ћ": "C", "Џ": "Dz", "Đ": "Dj",
	},
	"az": map[string]string{ /* Azerbaijani */
		"ç": "c", "ə": "e", "ğ": "g", "ı": "i", "ö": "o", "ş": "s", "ü": "u",
		"Ç": "C", "Ə": "E", "Ğ": "G", "İ": "I", "Ö": "O", "Ş": "S", "Ü": "U",
	},
}

var removeList = []string{
	"a", "an", "as", "at", "before", "but", "by", "for", "from",
	"is", "in", "into", "like", "of", "off", "on", "onto", "per",
	"since", "than", "the", "this", "that", "to", "up", "via",
	"with",
}

var removePattern = regexp.MustCompile("[^\\s_\\-a-zA-Z0-9]")

type parser struct {
	language   string
	maxLength  int
	text       string
	parsedText string
	removeList []string
}

// NewParser creates a new parser object
func NewParser() *parser {
	return &parser{}
}

func (p *parser) SetLanguage(language string) *parser {
	p.language = language
	p.parsedText = ""
	return p
}

func (p *parser) SetMaxLength(length int) *parser {
	p.maxLength = length
	p.parsedText = ""
	return p
}

func (p *parser) SetText(text string) *parser {
	p.parsedText = ""
	p.text = text
	return p
}

func (p *parser) AddToRemoveList(word string) *parser {
	p.removeList = append(p.removeList, strings.TrimSpace(word))
	return p
}

func (p *parser) RemoveFromRemoveList(word string) *parser {
	word = strings.TrimSpace(word)
	if len(p.removeList) > 0 {
		for i, w := range removeList {
			if w == word {
				p.removeList = append(p.removeList[:i], p.removeList[i+1:]...)
			}
		}
	}
	return p
}

func (p *parser) Parse() string {
	if p.parsedText != "" {
		return p.parsedText
	}

	text := strings.TrimSpace(p.text)

	if p.language != "" {
		chars, ok := charsMap[p.language]
		if ok {
			for from, to := range chars {
				text = strings.Replace(text, from, to, -1)
			}
		}
	}

	for lang, chars := range charsMap {
		// we did this replacement already
		if p.language != "" && p.language == lang {
			continue
		}
		for from, to := range chars {
			text = strings.Replace(text, from, to, -1)
		}
	}

	for _, word := range removeList {
		text = strings.Replace(text, " "+word, "", -1)
		text = strings.Replace(text, word+" ", "", -1)
	}

	if len(p.removeList) > 0 {
		for _, word := range p.removeList {
			text = strings.Replace(text, " "+word, "", -1)
			text = strings.Replace(text, word+" ", "", -1)
		}
	}

	text = removePattern.ReplaceAllString(text, "")

	text = strings.Replace(text, "_", " ", -1)
	text = strings.TrimSpace(text)
	text = strings.Replace(text, " ", "-", -1)
	text = strings.Replace(text, "--", "-", -1)
	text = strings.ToLower(text)

	if p.maxLength > 0 && len(text) > p.maxLength {
		text = text[0:p.maxLength]
	}

	p.parsedText = strings.Trim(text, "-")
	return p.parsedText
}
