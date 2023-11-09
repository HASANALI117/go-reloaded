package reboot

import "testing"

func TestNumberConversion(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "1E (hex) files were added.",
			expected: "30 files were added.",
		},
		{
			input:    "It has been 10 (bin) years",
			expected: "It has been 2 years",
		},
		{
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			// Test a valid hexadecimal number with the modifier correctly positioned.
			input:    "Convert 1F (hex) to decimal.",
			expected: "Convert 31 to decimal.",
		},
		{
			// Test a valid binary number with the modifier correctly positioned.
			input:    "Convert 1010 (bin) to decimal.",
			expected: "Convert 10 to decimal.",
		},
		{
			// Test mixed case and additional text with modifiers in the correct position.
			input:    "Mixed Case 1A (HEX) 1101 (bIn) Correctly Positioned.",
			expected: "Mixed Case 26 13 Correctly Positioned.",
		},
		{
			// Test multiple modifiers in the correct position.
			input:    "Multiple (hex) 1010 (bin) Convert Correctly Positioned.",
			expected: "Multiple (hex) 10 Convert Correctly Positioned.",
		},
	}

	for _, test := range tests {
		t.Run("TestAllConversions", func(t *testing.T) {
			result := NumberConversion(test.input)
			if result != test.expected {
				t.Errorf("Input:\n%s\nExpected:\n%s\nGot:\n%s", test.input, test.expected, result)
			}
		})
	}
}

func TestWordConversion(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "it (cap) was the best of times",
			expected: "It was the best of times",
		},
		{
			input:    "it was the worst of times (up)",
			expected: "it was the worst of TIMES",
		},
		{
			input:    "it was the age of foolishness don't (cap, 6)",
			expected: "It Was The Age Of Foolishness Don't",
		},
		{
			input:    "IT WAS THE (low, 0  ) winter of despair.",
			expected: "IT WAS THE (low, 0  ) winter of despair.",
		},
	}

	for _, test := range tests {
		t.Run("TestAllConversions", func(t *testing.T) {
			result := WordConversion(test.input)
			if result != test.expected {
				t.Errorf("Input:\n%s\nExpected:\n%s\nGot:\n%s", test.input, test.expected, result)
			}
		})
	}
}

func TestPunct2(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{
			input:    "As Elton John said: '    I am the most well-known 'person' in the world '",
			expected: "As Elton John said: 'I am the most well-known 'person' in the world'",
		},
	}

	for _, test := range tests {
		t.Run("TestAllConversions", func(t *testing.T) {
			result := Punct2(test.input)
			if result != test.expected {
				t.Errorf("Input:\n%s\nExpected:\n%s\nGot:\n%s", test.input, test.expected, result)
			}
		})
	}
}
