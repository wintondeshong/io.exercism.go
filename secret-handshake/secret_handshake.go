package secret

import (
	"strconv"
)

const testVersion = 1

var translations = []string{
	"wink",
	"double blink",
	"close your eyes",
	"jump",
}

// Public Methods
// --------------

func Handshake(code uint) (result []string) {
	s := strconv.FormatUint(uint64(code), 2)

	if s == "0" {
		return nil
	}

	for x := 1; x <= len(s); x++ {
		if s[len(s)-x] == '1' {
			result = append(result, translations[x-1])
		}
		if x == len(translations) {
			break
		}
	}

	if len(s) == 5 && s[0] == '1' {
		reverse(result)
	}

	return result
}

// Private Methods
// ---------------

func reverse(l []string) {
	for i := len(l)/2 - 1; i >= 0; i-- {
		opp := len(l) - 1 - i
		l[i], l[opp] = l[opp], l[i]
	}
}
