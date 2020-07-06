package dateutils

import "time"

const (
	apiDateLayout = "2006-01-02T15:04:05Z"
	apiDBLayout   = "2006-01-02 15:04:05"
)

// GetNow returns current UTC timestamp
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowString returns current timestamp
func GetNowString() string {
	return GetNow().Format(apiDateLayout)
}

// GetNowDBFormat returns DB layout time string
func GetNowDBFormat() string {
	return GetNow().Format(apiDBLayout)
}
