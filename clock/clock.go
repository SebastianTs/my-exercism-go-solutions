package clock

import (
	"fmt"
)

const testVersion = 4

// dayInMin is the duration of a day in Minutes
const dayInMin = 24 * 60

// Clock represents a time in minutes
type Clock int

// New creates a clock with the given time as int parameters
func New(hour, minute int) (c Clock) {

	return c.Add(hour*60 + minute)
}

// String prints the time a Clock represents
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c/60, c%60)
}

// Add adds and subtracts Time (Minutes) to a Clock
func (c Clock) Add(minutes int) Clock {
	c += Clock(minutes)
	c %= dayInMin //day
	if c < 0 {
		c += dayInMin
	}
	return c
}
