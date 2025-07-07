package valueobject

import "errors"

type Locale string

const (
	LocaleEN Locale = "en"
	LocaleRU Locale = "ru"
)

func NewLocale(val string) (Locale, error) {
	switch val {
	case "en", "ru":
		return Locale(val), nil
	default:
		return "", errors.New("invalid locale")
	}
}
