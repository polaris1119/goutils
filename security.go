package goutils

import (
	"net/url"
	"sort"
)

// GenSign 根据输入参数进行签名
func GenSign(args url.Values, secretKey string) string {
	keys := make([]string, 0, len(args))
	for k := range args {
		if k == "sign" {
			continue
		}
		keys = append(keys, k)
	}
	sort.Sort(sort.StringSlice(keys))

	buffer := NewBuffer()
	for _, k := range keys {
		buffer.Append(k).Append("=").Append(args.Get(k))
	}

	buffer.Append(secretKey)

	return Md5(buffer.String())
}
