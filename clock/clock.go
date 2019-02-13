package clock

import (
	"fmt"
	"time"
)

// Clock is a representation of a 24-hr clock
type Clock struct {
	H int
	M int
}

// New creates a new Clock with specified time in 24 hr format
func New(h, m int) Clock {
	h, m = normalize(h, m, 0)
	return Clock{
		H: h,
		M: m,
	}
}

// String returns the string representation of a clock
func (c Clock) String() string {
	return fmt.Sprintf("%02d:%02d", c.H, c.M)
}

// Add adds time to a clock
func (c Clock) Add(n int) Clock {
	c.H, c.M = normalize(c.H, c.M, n)
	return c
}

// Subtract subtracts time from a clock
func (c Clock) Subtract(n int) Clock {
	c.H, c.M = normalize(c.H, c.M, -n)
	return c
}

// normalize converts h and m into appropriate 24 hour format
func normalize(h, m, a int) (int, int) {

	// create a time object
	t, _ := time.Parse("15:04", "00:00")

	// convert total minutes to a duration
	d, _ := time.ParseDuration(fmt.Sprintf("%dm", (h*60)+m+a))

	// add duration to time object
	newTime := t.Add(d)

	// return 24 hour representation of hours and minutes
	return newTime.Hour(), newTime.Minute()
}
