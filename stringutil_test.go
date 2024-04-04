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

func TestCountSymbols(t *testing.T) {
    tests := []struct {
        input    string
        expected int
    }{
        {"hello", 5},
        {"12345", 0},
        {"hello 123", 5},
        {"こんにちは", 5}, // Japanese characters count as symbols
    }

    for _, tt := range tests {
        result := CountSymbols(tt.input)
        if result != tt.expected {
            t.Errorf("CountSymbols(%s) = %d; want %d", tt.input, result, tt.expected)
        }
    }
}