package log_model

import (
	"encoding/json"
	"fmt"
	"strings"
)

type Level int

const (
	ErrorLevel Level = 1 + iota
	WarnLevel
	InfoLevel
	DebugLevel
)

func (s Level) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}

func (s Level) String() string {
	var str string
	switch s {
	case DebugLevel:
		str = "debug"
	case InfoLevel:
		str = "info"
	case WarnLevel:
		str = "warning"
	case ErrorLevel:
		str = "error"
	}

	return str
}

// ParseLevel 字符串解析为level
func ParseLevel(lvl string) (Level, error) {
	switch strings.ToLower(lvl) {
	case "error":
		return ErrorLevel, nil
	case "warn", "warning":
		return WarnLevel, nil
	case "info":
		return InfoLevel, nil
	case "debug":
		return DebugLevel, nil
	}

	var l Level
	return l, fmt.Errorf("not a valid logrus Level: %q", lvl)
}
