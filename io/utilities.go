package io

import (
	"bytes"
	"compress/gzip"
	"io"
)

func GzipEncode(input []byte) ([]byte, error) {
	var buf bytes.Buffer
	w := gzip.NewWriter(&buf)
	defer w.Close()
	_, err := w.Write(input)
	if err == nil {
		w.Close()
	}
	return buf.Bytes(), err
}

func GzipDecode(input []byte) ([]byte, error) {
	reader, err := gzip.NewReader(bytes.NewBuffer(input))
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	var result []byte
	buf := make([]byte, 1024)
	for {
		count, err := reader.Read(buf)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if count > 0 {
			result = append(result, buf[0:count]...)
		} else {
			break
		}
	}
	return result, nil
}
