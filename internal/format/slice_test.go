package format

import (
	"reflect"
	"testing"
)

func TestUniqueSlice(t *testing.T) {
	tests := []struct {
		name     string
		slice    []string
		expected []string
	}{
		{
			name:     "empty slice",
			slice:    []string{},
			expected: []string{},
		},
		{
			name:     "single item",
			slice:    []string{"a"},
			expected: []string{"a"},
		},
		{
			name:     "multiple items",
			slice:    []string{"a", "b", "a", "c", "b"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "multiple items, no duplicates",
			slice:    []string{"a", "b", "c"},
			expected: []string{"a", "b", "c"},
		},
		{
			name:     "multiple items, all duplicates",
			slice:    []string{"a", "a", "a"},
			expected: []string{"a"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UniqueSlice(test.slice)

			if len(result) != len(test.expected) {
				t.Errorf("Expected %d items, got %d", len(test.expected), len(result))
				return
			}

			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}

		})
	}
}
