package acronym

import (
	"fmt"
	"strings"
	"unicode"
)

const testVersion = 1

func abbreviate(input string) (result string) {
	var prev rune

	for i, r := range input {
		if i == 0 || prev == ' ' || prev == '-' || (unicode.IsUpper(r) && unicode.IsLower(prev)) {
			result += fmt.Sprintf("%c", r)
		}
		prev = r
	}

	return strings.ToUpper(result)
}
