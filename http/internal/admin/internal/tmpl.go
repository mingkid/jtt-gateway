package internal

import "time"

// FormatAsTimestamp 时间格式化
func FormatAsTimestamp(ts int64) string {
	if ts == 0 {
		return "-"
	}
	return time.Unix(ts, 0).Format("2006/01/02 03:04:05")
}
