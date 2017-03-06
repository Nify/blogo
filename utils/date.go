package utils

import "time"

const (
	timeformat = "2006-01-02"
)

// ParseDate parses a formatted string and returns the time value it represents.
// The layout  defines the format by showing how the reference time.
func ParseDate(value string) (time.Time, error) {
	date, err := time.Parse(timeformat, value)
	if err != nil {
		date, err = time.ParseInLocation(timeformat, value, time.Now().Location())
		if err != nil {
			return time.Now(), err
		}
	}
	return date, nil
}
