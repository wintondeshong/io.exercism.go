package bob

import (
	"strings"
	"unicode"
)

const testVersion = 2

// Public Methods
// --------------

// Hey responds to the supplied input depending upon its format.
// Input anything outside of a valid question and you'll get some lip.
func Hey(s string) string {
	s = strings.TrimSpace(s)

	if s == "" {
		return "Fine. Be that way!"
	}

	if isAllCaps(s) {
		return "Whoa, chill out!"
	}

	if strings.HasSuffix(s, "?") {
		return "Sure."
	}

	return "Whatever."
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
