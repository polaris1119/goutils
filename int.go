// Copyright 2016 polaris. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	polaris@studygolang.com

package goutils

import (
	"log"
	"strconv"
)

// MustInt 字符串转int
func MustInt(s string, defaultVal ...int) int {
	getDefault := func() int {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	if s == "" {
		return getDefault()
	}

	i, err := strconv.Atoi(s)
	if err != nil {
		log.Println("goutils MustInt strconv.Atoi error:", err)
		return getDefault()
	}

	return i
}

// MustInt 字符串转int64
func MustInt64(s string, defaultVal ...int64) int64 {
	getDefault := func() int64 {
		if len(defaultVal) > 0 {
			return defaultVal[0]
		}
		return 0
	}

	if s == "" {
		return getDefault()
	}

	i, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		log.Println("goutils MustInt64 strconv.ParseInt error:", err)
		return getDefault()
	}

	return i
}
