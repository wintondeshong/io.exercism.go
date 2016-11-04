package clock

import (
	"fmt"
	"math"
)

const (
	hoursPerDay    = 24
	minutesPerDay  = 1440
	minutesPerHour = 60
	testVersion    = 4
)

// Clock type containing only hour and minute precision
type Clock struct {
	hour   int
	minute int
}

// New clock constructor
func New(hour, minute int) Clock {
	return Clock{}.Add(hour*minutesPerHour + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

// Add offsets the current clock by the provided positive or negative minutes integer
func (c Clock) Add(minutes int) Clock {
	// normalize minutes to fraction of single day
	if minutes > minutesPerDay || minutes < -minutesPerDay {
		minutes %= minutesPerDay
	}

	// flatten current time to minutes and add incoming minutes
	var totalMinutes = c.hour*minutesPerHour + c.minute + minutes

	// if greater than +/- minutesPerDay, remove a day's worth of minutes
	if totalMinutes > minutesPerDay {
		totalMinutes -= minutesPerDay
	} else if totalMinutes < -minutesPerDay {
		totalMinutes += minutesPerDay
	}

	if totalMinutes < 0 {
		totalMinutes = minutesPerDay + totalMinutes
	}

	// convert to hours and minutes
	var h = int(math.Floor(float64(totalMinutes) / minutesPerHour))
	var m = totalMinutes - minutesPerHour*h

	// max hours per day? tip over to new day
	if h >= hoursPerDay {
		h -= hoursPerDay
	}

	return Clock{h, m}
}
