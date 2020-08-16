package palindrome

import "testing"

var tests = []struct {
	input    string
	expected bool
}{
	{
		"racecar",
		true,
	},
	{
		"cac",
		true,
	},
	{
		"caac",
		true,
	},
	{
		"notracecar",
		false,
	},
	{
		"",
		true,
	},
	{
		"a",
		true,
	},
}

func TestPalindrome(t *testing.T) {
	for _, tc := range tests {
		actual := IsPalindrome(tc.input)
		if actual != tc.expected {
			t.Fatalf("Input %s failed! Expected : %t, actual: %t", tc.input, tc.expected, actual)
		}
	}
}

func BenchmarkPalindrome(b *testing.B) {

	b.ResetTimer()
	for _, tc := range tests {
		IsPalindrome(tc.input)

	}
}
