package goutils

import (
	"log"
	"strconv"
	"strings"
)

func getFloatDefault(defaultVals ...float64) float64 {
	if len(defaultVals) > 0 {
		return defaultVals[0]
	}
	return 0.0
}

func mustFloat(s string, defaultVals ...float64) float64 {
	if s == "" {
		return getFloatDefault()
	}

	f, err := strconv.ParseFloat(strings.TrimSpace(s), 64)
	if err != nil {
		log.Println("goutils MustFloat strconv.ParseFloat error:", err)
		return getFloatDefault()
	}

	return f
}

func MustFloat(inter interface{}, defaultVals ...float64) float64 {
	switch v := inter.(type) {
	case float64:
		return v
	case string:
		return mustFloat(v, defaultVals...)
	case int64:
		return float64(v)
	case float32:
		return float64(v)
	default:
		return getFloatDefault(defaultVals...)
	}
}
