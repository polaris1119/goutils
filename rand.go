package goutils

import (
	"fmt"
	"math/rand"
	"strconv"

	"github.com/jmcvetta/randutil"
)

const DefaultCharset = randutil.Alphanumeric + "!%@#"

// RandString 产生随机字符串，可用于密码等
func RandString(n int, defaultCharsets ...string) string {
	charset := DefaultCharset
	if len(defaultCharsets) > 0 {
		charset = defaultCharsets[0]
	}

	result, err := randutil.String(n, charset)
	if err != nil {
		fmt.Println("goutils RandString error:", err)
		return Md5(strconv.Itoa(rand.Intn(999999)))
	}

	return result
}
