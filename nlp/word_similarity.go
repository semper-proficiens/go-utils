package nlp

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"log"
)

// KeyValue represents a key-value pair.
type KeyValue struct {
	Key   string
	Value string
}

// CalculateSimilarity calculates the similarity score between two strings using Levenshtein distance.
func CalculateSimilarity(a, b string) float64 {
	distance := levenshtein.DistanceForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	maxLen := max(len(a), len(b))
	if maxLen == 0 {
		return 1.0
	}
	return 1.0 - float64(distance)/float64(maxLen)
}

// RemoveDuplicateKV removes duplicate key-value pairs based on a similarity threshold. This assumes all Keys in the
// the KeyValue object are identical, so we only focus on the similarity of Value.
func RemoveDuplicateKV(kv []KeyValue, threshold float64) []KeyValue {
	var uniqueKV []KeyValue

	for i, kvPair := range kv {
		isDuplicate := false
		for j := 0; j < i; j++ {
			similarity := CalculateSimilarity(kvPair.Value, kv[j].Value)
			if similarity >= threshold {
				log.Printf("Excluded potential duplicate value \"%s\" and \"%s\" with similarity score %.2f", kvPair.Value, kv[j].Value, similarity)
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			uniqueKV = append(uniqueKV, kvPair)
		}
	}

	return uniqueKV
}
