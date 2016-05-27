// Copyright 2016 polaris. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author: polaris	polaris@studygolang.com

package goutils

import (
	"fmt"
	"log"
	"runtime"
	"strconv"
	"strings"
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

	i, err := strconv.Atoi(strings.TrimSpace(s))
	if err != nil {
		msg := "goutils MustInt strconv.Atoi error:" + err.Error()
		// 加上文件调用和行号
		_, callerFile, line, ok := runtime.Caller(1)
		if ok {
			msg += fmt.Sprintf("file:%s,line:%d", callerFile, line)
		}
		log.Println(msg)
		return getDefault()
	}

	return i
}

// MustInt64 字符串转int64
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

	i, err := strconv.ParseInt(strings.TrimSpace(s), 10, 64)
	if err != nil {
		msg := "goutils MustInt64 strconv.ParseInt error:" + err.Error()
		// 加上文件调用和行号
		_, callerFile, line, ok := runtime.Caller(1)
		if ok {
			msg += fmt.Sprintf("file:%s,line:%d", callerFile, line)
		}
		log.Println(msg)
		log.Println("goutils MustInt64 strconv.ParseInt error:", err)
		return getDefault()
	}

	return i
}
