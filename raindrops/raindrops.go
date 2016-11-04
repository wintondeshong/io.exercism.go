package raindrops

import (
	"fmt"
	"strconv"
)

const testVersion = 2

var factorMatches = []struct {
	factor    int
	character string
}{
	{3, "i"},
	{5, "a"},
	{7, "o"},
}

func Convert(i int) (result string) {
	for _, match := range factorMatches {
		if i%match.factor == 0 {
			result += fmt.Sprintf("Pl%sng", match.character)
		}
	}

	if len(result) == 0 {
		return strconv.Itoa(i)
	}

	return
}
