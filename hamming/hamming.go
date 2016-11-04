package hamming

import "errors"

const testVersion = 5

var ErrDifferentParamLengths = errors.New("Distance: parameters must be the same length")

// Distance calculates the hamming distance between two DNA strands [0-N].
// If the supplied strings are not of the same length, returns -1.
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, ErrDifferentParamLengths
	}

	diff := 0

	for i, _ := range a {
		if a[i] != b[i] {
			diff++
		}
	}

	return diff, nil
}
