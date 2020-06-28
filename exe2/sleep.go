package exe2

import "runtime"

// SleepCommand returns the appropriate sleep exec command for the
// current platform. This is here for testing purposes.
//
// Windows Batch does not support sleep. Use powershell sleep
// there instead.
func SleepCommand(sleepTime string) string {
	if runtime.GOOS == "windows" {
		return "powershell sleep " + sleepTime
	}

	return "sleep " + sleepTime
}
