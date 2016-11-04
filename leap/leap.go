package leap

const testVersion = 2

// IsLeapYear returns a boolean value indicating whether the year provided
// is a leap year according to the Gregorian calendar.
func IsLeapYear(year int) bool {
	if year%4 != 0 {
		return false
	}

	return year%100 != 0 || year%400 == 0
}
