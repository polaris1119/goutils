// Copyright 2016 polaris. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author: polaris	polaris@studygolang.com

package goutils

import (
	"crypto/md5"
	"encoding/base64"
	"fmt"
	"io"
)

func Md5(text string) string {
	hashMd5 := md5.New()
	io.WriteString(hashMd5, text)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Md5Buf(buf []byte) string {
	hashMd5 := md5.New()
	hashMd5.Write(buf)
	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Md5File(reader io.Reader) string {
	var buf = make([]byte, 4096)
	hashMd5 := md5.New()
	for {
		n, err := reader.Read(buf)
		if err == io.EOF && n == 0 {
			break
		}
		if err != nil && err != io.EOF {
			break
		}

		hashMd5.Write(buf[:n])
	}

	return fmt.Sprintf("%x", hashMd5.Sum(nil))
}

func Base64Encode(data string) string {
	return base64.URLEncoding.EncodeToString([]byte(data))
}

func Base64Decode(data string) string {
	b, err := base64.URLEncoding.DecodeString(data)
	if err != nil {
		return ""
	}
	return string(b)
}
