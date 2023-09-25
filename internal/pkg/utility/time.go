package utility

import (
	"time"
)

// GetNow returns the current time in UTC
func GetNow() time.Time {
	return time.Now().UTC()
}

// GetNowUnix returns the current time in UTC as Unix time
func GetNowUnix() int64 {
	return GetNow().Unix()
}
