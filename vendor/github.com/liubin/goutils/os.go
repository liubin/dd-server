package goutils

import (
	"os"
	"strconv"
)

func GetEnvInt64(key string, defaultValue int64) int64 {
	s := os.Getenv(key)
	if v, e := strconv.ParseInt(s, 10, 64); e == nil {
		return v
	} else {
		return defaultValue
	}
}

func GetEnvString(key string, defaultValue string) string {
	if s := os.Getenv(key); s != "" {
		return s
	} else {
		return defaultValue
	}
}
