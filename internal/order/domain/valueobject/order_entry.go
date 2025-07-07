package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

type OrderEntry string

var entryRe = regexp.MustCompile(`^[A-Z0-9]+$`)

func NewOrderEntry(val string) (OrderEntry, error) {
	val = strings.ToUpper(val)
	if val == "" || !entryRe.MatchString(val) {
		return "", errors.New("invalid order entry")
	}
	return OrderEntry(val), nil
}
