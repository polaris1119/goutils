package goutils_test

import (
	"testing"

	"github.com/polaris1119/goutils"
)

func TestConvertString(t *testing.T) {
	tests := []struct {
		input    interface{}
		expected string
		prec     int
	}{
		{"abc", "abc", 0},
		{3.20, "3.2", 1},
		{3.20, "3.20", 2},
		{3.21, "3.21", 2},
		{float32(3.21), "3.21", 2},
		{1, "1", 0},
		{uint(1), "1", 0},
		{int8(1), "1", 0},
		{uint8(1), "1", 0},
		{false, "false", 0},
	}

	for _, test := range tests {
		actual := goutils.ConvertString(test.input, test.prec)
		if test.expected != actual {
			t.Errorf("ConvertString(%q, %q)=%s, expected:%s", test.input, test.prec, actual, test.expected)
		}
	}
}
