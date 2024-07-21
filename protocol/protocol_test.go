package protocol

import "testing"

func TestProcessMessage(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"HELLO", "WORLD"},
		{"hello", "UNKNOWN COMMAND"},
		{"FOO", "UNKNOWN COMMAND"},
	}

	for _, test := range tests {
		result := ProcessMessage(test.input)
		if result != test.expected {
			t.Errorf("For input '%s', expected '%s', but got '%s'", test.input, test.expected, result)
		}
	}
}
