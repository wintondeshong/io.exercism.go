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

type Clock int

func New(hour, minute int) Clock {
	return Clock{}.Add(hour*minutesPerHour + minute)
}

func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.hour, c.minute)
}

func (c Clock) Add(minutes int) Clock {
	h := (c / 60) % 24
	m := c % 60
	return New(int(h), int(m)+minutes)
}
