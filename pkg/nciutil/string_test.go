package nciutil

import (
	"testing"
)

func TestFirstNonEmpty(t *testing.T) {
	tests := []struct {
		strings []string
		result  string
	}{
		{[]string{"", "hello", "world"}, "hello"},
		{[]string{"", "", "hello"}, "hello"},
		{[]string{"", "", ""}, ""},
		{[]string{"hello", "world", "golang"}, "hello"},
		{[]string{"golang", "world", "hello"}, "golang"},
	}

	for _, test := range tests {
		res := FirstNonEmpty(test.strings)
		if res != test.result {
			t.Errorf("FirstNonEmpty(%v) = %v, want %v", test.strings, res, test.result)
		}
	}
}

func TestGetValueFromMap(t *testing.T) {
	testCases := []struct {
		name     string
		m        map[string]string
		key      string
		expected string
	}{
		{
			name:     "Key present in map",
			m:        map[string]string{"foo": "bar"},
			key:      "foo",
			expected: "bar",
		},
		{
			name:     "Key not present in map",
			m:        map[string]string{"foo": "bar"},
			key:      "baz",
			expected: "",
		},
		{
			name:     "Empty map",
			m:        map[string]string{},
			key:      "foo",
			expected: "",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			result := GetValueFromMap(tc.m, tc.key)
			if result != tc.expected {
				t.Errorf("Expected %q but got %q", tc.expected, result)
			}
		})
	}
}
