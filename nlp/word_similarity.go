package nlp

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// CalculateSimilarity calculates the similarity score between two strings using Levenshtein distance.
func CalculateSimilarity(a, b string) float64 {
	distance := levenshtein.DistanceForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	maxLen := max(len(a), len(b))
	if maxLen == 0 {
		return 1.0
	}
	return 1.0 - float64(distance)/float64(maxLen)
}
