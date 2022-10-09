package util

import "time"

// timestamps are expressed in number of seconds since midnight 1/1/1980 (I think...)
var apolloEpoch = time.Date(1980, time.January, 1, 0, 0, 0, 0, time.UTC)

func FormatTimestamp(timestamp int32) string {
	return apolloEpoch.Add(time.Duration(timestamp) * time.Second).Format(time.RFC1123)
}
