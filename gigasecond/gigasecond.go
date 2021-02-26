package gigasecond

import "time"

// AddGigasecond adds 10^9 seconds to the given time, t.
func AddGigasecond(t time.Time) time.Time {
	return t.Add(1e9 * time.Second)
}
