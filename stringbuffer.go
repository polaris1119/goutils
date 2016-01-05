// Copyright 2016 polaris. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.
// http://studygolang.com
// Author：polaris	polaris@studygolang.com

package goutils

import (
	"bytes"
	"log"
	"strconv"
)

// 内嵌bytes.Buffer，支持连写
type Buffer struct {
	*bytes.Buffer
}

func NewBuffer() *Buffer {
	return &Buffer{Buffer: new(bytes.Buffer)}
}

func (b *Buffer) Append(s string) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()

	b.WriteString(s)
	return b
}

func (b *Buffer) AppendInt(i int64) *Buffer {
	return b.Append(strconv.FormatInt(i, 10))
}

func (b *Buffer) AppendUint(i uint64) *Buffer {
	return b.Append(strconv.FormatUint(i, 10))
}

func (b *Buffer) AppendBytes(p []byte) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()

	b.Write(p)

	return b
}

func (b *Buffer) AppendRune(r rune) *Buffer {
	defer func() {
		if err := recover(); err != nil {
			log.Println("*****内存不够了！******")
		}
	}()

	b.WriteRune(r)

	return b
}
