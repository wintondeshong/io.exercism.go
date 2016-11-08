package bob

import (
	"strings"
	"unicode"
)

const (
	testVersion        = 2
	ResponseToQuestion = "Sure."
	ResponseToYelling  = "Whoa, chill out!"
	ResponseToSilence  = "Fine. Be that way!"
	ResponseDefault    = "Whatever."
)

// Public Methods
// --------------

// Hey responds to the supplied input depending upon its format.
// Input anything outside of a valid question and you'll get some lip.
func Hey(s string) string {
	s = strings.TrimSpace(s)

	if s == "" {
		return ResponseToSilence
	}

	if isAllCaps(s) {
		return ResponseToYelling
	}

	if strings.HasSuffix(s, "?") {
		return ResponseToQuestion
	}

	return ResponseDefault
}

// Private Methods
// ---------------

func isAllCaps(s string) bool {
	hasLetters := false

	for _, c := range s {
		if unicode.IsLetter(c) {
			hasLetters = true

			if unicode.IsLower(c) {
				return false
			}
		}
	}

	return hasLetters
}
