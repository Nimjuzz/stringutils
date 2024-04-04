package stringutil

import "testing"

func TestReverse(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"12345", "54321"},
		{"", ""},
	}

	for _, tt := range tests {
		result := Reverse(tt.input)
		if result != tt.expected {
			t.Errorf("Reverse(%s) = %s; want %s", tt.input, result, tt.expected)
		}
	}
}