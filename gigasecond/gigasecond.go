package gigasecond

import (
	"time"
)

const testVersion = 4

// AddGigasecond return a given time with one Gigasecond added to it
func AddGigasecond(t time.Time) time.Time {

	return t.Add(1e9 * time.Second)
}
