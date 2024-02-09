package request

import (
	"errors"
	"io"
)

var ERR_CONTENT_TO_LARGE = errors.New("err content to large")
var ERR_NEGATIVE_SIZE = errors.New("err negative size")

func ReadLimitedBytes(r io.Reader, limit int64) ([]byte, error) {
	buf := make([]byte, limit)

	n, err := io.ReadFull(r, buf)

	if err != nil && err != io.EOF && err != io.ErrUnexpectedEOF {
		return nil, err
	}

	return buf[:n], nil
}

func ReadLimitedBytesWithSizeCheck(r io.Reader, limit int64) ([]byte, error) {
	buf, err := ReadLimitedBytes(r, limit+1)

	if err != nil {
		return nil, err
	}

	if len(buf) > int(limit) {
		return buf[:len(buf)-1], ERR_CONTENT_TO_LARGE
	}

	return buf, nil
}

var size_kib int64 = 1024
var size_mib int64 = 1024 * size_kib
var size_gib int64 = 1024 * size_mib

func CalculateKiB(size int64) (int64, error) {
	return genericCalculate(size, size_kib)
}

func CalculateMiB(size int64) (int64, error) {
	return genericCalculate(size, size_mib)
}

func CalculateGiB(size int64) (int64, error) {
	return genericCalculate(size, size_gib)
}

func genericCalculate(size int64, factor int64) (int64, error) {
	if size < 0 {
		return 0, ERR_NEGATIVE_SIZE
	}

	return size * factor, nil
}
