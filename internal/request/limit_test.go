package request_test

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/fschuermeyer/GoWordlytics/internal/request"
)

func TestReadLimitedBytes(t *testing.T) {
	tests := []struct {
		name        string
		reader      io.Reader
		limit       int64
		expected    []byte
		expectedErr error
	}{
		{
			name:        "read full bytes",
			reader:      bytes.NewBufferString("Hello, World!"),
			limit:       13,
			expected:    []byte("Hello, World!"),
			expectedErr: nil,
		},
		{
			name:        "read partial bytes",
			reader:      bytes.NewBufferString("Hello, World!"),
			limit:       5,
			expected:    []byte("Hello"),
			expectedErr: nil,
		},
		{
			name:        "read empty bytes",
			reader:      bytes.NewBufferString(""),
			limit:       10,
			expected:    []byte{},
			expectedErr: nil,
		},
		{
			name:        "read error",
			reader:      &mockReader{err: errors.New("read error")},
			limit:       10,
			expected:    nil,
			expectedErr: errors.New("read error"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := request.ReadLimitedBytes(tt.reader, tt.limit)

			if !bytes.Equal(result, tt.expected) {
				t.Errorf("ReadLimitedBytes(%v, %v) = %v, want %v", tt.reader, tt.limit, result, tt.expected)
			}

			if (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr == nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("ReadLimitedBytes(%v, %v) error = %v, want %v", tt.reader, tt.limit, err, tt.expectedErr)
			}
		})
	}
}

type mockReader struct {
	err error
}

func (r *mockReader) Read(p []byte) (n int, err error) {
	return 0, r.err
}
func TestReadLimitedBytesWithSizeCheck(t *testing.T) {
	tests := []struct {
		name        string
		reader      io.Reader
		limit       int64
		expected    []byte
		expectedErr error
	}{
		{
			name:        "read full bytes",
			reader:      bytes.NewBufferString("Hello, World!"),
			limit:       13,
			expected:    []byte("Hello, World!"),
			expectedErr: nil,
		},
		{
			name:        "read partial bytes",
			reader:      bytes.NewBufferString("Hello, World!"),
			limit:       5,
			expected:    []byte("Hello"),
			expectedErr: request.ERR_CONTENT_TO_LARGE,
		},
		{
			name:        "read empty bytes",
			reader:      bytes.NewBufferString(""),
			limit:       10,
			expected:    []byte{},
			expectedErr: nil,
		},
		{
			name:        "read error",
			reader:      &mockReader{err: errors.New("read error")},
			limit:       10,
			expected:    nil,
			expectedErr: errors.New("read error"),
		},
		{
			name:        "read bytes exceeding limit",
			reader:      bytes.NewBufferString("Hello, World!"),
			limit:       5,
			expected:    []byte("Hello"),
			expectedErr: request.ERR_CONTENT_TO_LARGE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := request.ReadLimitedBytesWithSizeCheck(tt.reader, tt.limit)

			if !bytes.Equal(result, tt.expected) {
				t.Errorf("ReadLimitedBytesWithSizeCheck(%v, %v) = %v, want %v", tt.reader, tt.limit, result, tt.expected)
			}

			if (err == nil && tt.expectedErr != nil) || (err != nil && tt.expectedErr == nil) || (err != nil && tt.expectedErr != nil && err.Error() != tt.expectedErr.Error()) {
				t.Errorf("ReadLimitedBytesWithSizeCheck(%v, %v) error = %v, want %v", tt.reader, tt.limit, err, tt.expectedErr)
			}
		})
	}
}

func TestCalculateKiB(t *testing.T) {
	tests := []struct {
		name        string
		size        int64
		expected    int64
		expectedErr error
	}{
		{
			name:     "positive size",
			size:     10,
			expected: 10240,
		},
		{
			name:     "zero size",
			size:     0,
			expected: 0,
		},
		{
			name:        "negative size",
			size:        -5,
			expected:    0,
			expectedErr: request.ERR_NEGATIVE_SIZE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := request.CalculateKiB(tt.size)

			if err != nil && tt.expectedErr != err {
				t.Errorf("CalculateKiB(%v) error = %v, want %v", tt.size, err, tt.expectedErr)
			}

			if result != tt.expected {
				t.Errorf("CalculateKiB(%v) = %v, want %v", tt.size, result, tt.expected)
			}
		})
	}
}
