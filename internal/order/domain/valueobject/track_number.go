package valueobject

import (
	"errors"
	"regexp"
	"strings"
)

type TrackNumber string

var trackRe = regexp.MustCompile(`^[A-Z0-9]+$`)

func NewTrackNumber(val string) (TrackNumber, error) {
	val = strings.ToUpper(val)
	if val == "" {
		return "", errors.New("track number must not be empty")
	}
	if !trackRe.MatchString(val) {
		return "", errors.New("track number must be upper-case A-Z0-9")
	}
	return TrackNumber(val), nil
}
