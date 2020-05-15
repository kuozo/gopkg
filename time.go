package gopkg

import "time"

// FormatNow  format now
func FormatNow() {
	time.Now().Format(time.RFC850)
}
