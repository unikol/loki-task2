package logger

import (
	"os"
	"path/filepath"
)

func CreateLogger(logfilePath string) (*os.File, error) {
	logDir := filepath.Dir(logfilePath)
	if err := os.MkdirAll(logDir, os.ModePerm); err != nil {
		return nil, err
	}

	logFile, err := os.OpenFile(logfilePath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	return logFile, nil
}
