package nlp

import "testing"

// TestCalculateSimilarity tests the CalculateSimilarity function.
func TestCalculateSimilarity(t *testing.T) {
	tests := []struct {
		a, b     string
		expected float64
	}{
		{"kitten", "sitting", 0.2857142857142857},
		{"flaw", "lawn", 0.5},
		{"", "", 1.0},
		{"", "nonempty", 0.0},
		{"nonempty", "", 0.0},
		{"identical", "identical", 1.0},
		{"abc", "abc", 1.0},
		{"abc", "def", -1.0},
		{"abc", "abcd", 0.75},
	}

	for _, test := range tests {
		t.Run(test.a+"_"+test.b, func(t *testing.T) {
			result := CalculateSimilarity(test.a, test.b)
			if result != test.expected {
				t.Errorf("CalculateSimilarity(%q, %q) = %v; want %v", test.a, test.b, result, test.expected)
			}
		})
	}
}
