package util

import "time"

const layout = "2006-01-02T15:04:05"

func StringToTime(value string) time.Time {
	result, err := time.Parse(layout, value)

	if err != nil {
		return time.Time{}
	}

	return result
}
