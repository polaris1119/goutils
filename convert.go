// Copyright 2016 polaris. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author: polaris	polaris@studygolang.com

package goutils

import (
	"fmt"
	"strconv"
)

func ConvertString(inter interface{}, precs ...int) string {
	switch v := inter.(type) {
	case string:
		return v
	case float64:
		prec := 0
		if len(precs) > 0 {
			prec = precs[0]
		}
		return strconv.FormatFloat(v, 'f', prec, 64)
	case int64:
		return strconv.FormatInt(v, 10)
	case uint64:
		return strconv.FormatUint(v, 10)
	case int:
		return strconv.Itoa(v)
	case uint:
		return strconv.FormatUint(uint64(v), 10)
	default:
		return fmt.Sprintf("%v", inter)
	}
}
