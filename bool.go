package goutils

import (
	"log"
	"strconv"
	"strings"
)

func MustBool(s string, defaultVal ...bool) bool {
	getDefault := func() bool {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return false
	}

	if s == "" {
		return getDefault()
	}

	b, err := strconv.ParseBool(strings.TrimSpace(s))
	if err != nil {
		log.Println("goutils MustBool strconv.ParseBool error:", err)
		return getDefault()
	}

	return b
}
