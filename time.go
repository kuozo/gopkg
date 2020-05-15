package gopkg

import "time"

// FormatNow  format now
func FormatNow() string {
	return time.Now().Format(time.RFC850)
}
