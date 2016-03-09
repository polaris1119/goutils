package goutils

import (
	"log"
	"strconv"
)

func MustFloat(s string, defaultVal ...float64) float64 {
	getDefault := func() float64 {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0.0
	}

	if s == "" {
		return getDefault()
	}

	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		log.Println("goutils MustFloat strconv.ParseFloat error:", err)
		return getDefault()
	}

	return f
}
