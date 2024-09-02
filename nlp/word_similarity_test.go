package nlp

import (
	"testing"
)

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

func TestRemoveDuplicateKV(t *testing.T) {
	// Test data
	kv := []KeyValue{
		{Key: "Title", Value: "Title Name1 Something"},
		{Key: "Title", Value: "Title Name1 Something Something"},
		{Key: "Title", Value: "Title Name2 Something"},
		{Key: "Title", Value: "Title Name2 Something Something"},
	}

	threshold := 0.6
	expected := []KeyValue{
		{Key: "Title", Value: "Title Name1 Something"},
	}

	// Call the function
	result := RemoveDuplicateKV(kv, threshold)

	// Check the result
	if len(result) != len(expected) {
		t.Errorf("Expected %d unique key-value pairs, got %d", len(expected), len(result))
	}

	for i, kvPair := range expected {
		if result[i].Key != kvPair.Key || result[i].Value != kvPair.Value {
			t.Errorf("Expected key-value pair {\"%s\":\"%s\"}, but got {\"%s\":\"%s\"}", kvPair.Key, kvPair.Value, result[i].Key, result[i].Value)
		}
	}
}
