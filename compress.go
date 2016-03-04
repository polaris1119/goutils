package goutils

import (
	"bytes"
	"compress/gzip"
	"errors"
	"io"
	"io/ioutil"
	"log"
	"sync"
)

var gzipWriterPool = sync.Pool{
	New: func() interface{} {
		return gzip.NewWriter(ioutil.Discard)
	},
}

var gzipHeader = []byte{
	0x1f, 0x8b, 0x08, 0x08, 0xf7, 0x5e, 0x14, 0x4a,
	0x00, 0x03, 0x65, 0x6d, 0x70, 0x74, 0x79, 0x2e,
	0x74, 0x78, 0x74, 0x00, 0x03, 0x00, 0x00, 0x00,
	0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
}
var gzipReaderPool = sync.Pool{
	New: func() interface{} {
		reader, err := gzip.NewReader(bytes.NewReader(gzipHeader))
		if err != nil {
			log.Println("gzip NewReader error:", err)
		}
		return reader
	},
}

func Gzip(b []byte, w io.Writer) (int, error) {
	writer := gzipWriterPool.Get().(*gzip.Writer)
	writer.Reset(w)

	defer func() {
		writer.Close()
		gzipWriterPool.Put(writer)
	}()

	return writer.Write(b)
}

func Gunzip(r io.Reader) ([]byte, error) {
	obj := gzipReaderPool.Get()
	if obj == nil {
		return nil, errors.New("gizp NewReader error")
	}

	reader := obj.(*gzip.Reader)
	reader.Reset(r)

	defer func() {
		reader.Close()
		gzipReaderPool.Put(reader)
	}()

	// 直接 reader.Read()，当数据量大时，读取不完整（需要循环）
	return ioutil.ReadAll(reader)
}
