package xgzip

import (
	"bytes"
	"compress/gzip"
	"io/ioutil"
)

func Decompress(input []byte) (string, error) {
	buf := bytes.NewBuffer(input)
	reader, err := gzip.NewReader(buf)
	if err != nil {
		return "", err
	}
	defer func() { _ = reader.Close() }()

	result, err := ioutil.ReadAll(reader)
	if err != nil {
		return "", err
	}

	return string(result), nil
}

func Compress(input string) ([]byte, error) {
	var buf bytes.Buffer
	gz := gzip.NewWriter(&buf)

	_, err := gz.Write([]byte(input))
	if err != nil {
		return nil, err
	}

	err = gz.Flush()
	if err != nil {
		return nil, err
	}

	err = gz.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}
