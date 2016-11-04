package gigasecond

import "time"

const (
	gigasecond  time.Duration = time.Second * 1e9
	testVersion               = 4
)

// AddGigasecond increases the supplied time by one Gigasecond
func AddGigasecond(t time.Time) time.Time {
	return t.Add(gigasecond)
}
