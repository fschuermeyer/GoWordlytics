package format

import "testing"

func TestInsertThousandSeparator(t *testing.T) {
	tests := []struct {
		number    int
		separator rune
		expected  string
	}{
		{1000, ',', "1,000"},
		{123456789, '.', "123.456.789"},
		{9876543210, '-', "9-876-543-210"},
		{999, '.', "999"},
		{0, ':', "0"},
	}

	for _, test := range tests {
		result := InsertThousandSeparator(test.number, test.separator)
		if result != test.expected {
			t.Errorf("InsertThousandSeparator(%d, '%c') = %s, expected %s", test.number, test.separator, result, test.expected)
		}
	}
}
