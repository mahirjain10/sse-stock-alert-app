package helpers

import "runtime"

func GetFileName() string {
	_, file, _, ok := runtime.Caller(2)
	if !ok {
		return "unknown"
	}
	return file
}

// getLineNumber retrieves the current line number
func GetLineNumber() int {
	_, _, line, ok := runtime.Caller(2)
	if !ok {
		return 0
	}
	return line
}