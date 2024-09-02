package nlp

import (
	"github.com/texttheater/golang-levenshtein/levenshtein"
	"log"
	"reflect"
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

// RemoveDuplicates removes duplicate items from a slice based on a similarity threshold.
// The fieldName parameter specifies the field to be used for similarity comparison.
// e.g.
//
//	articles := []MockArticle{
//			{Title: "Title Name1 Something"},
//			{Title: "Title Name1 Something Something"},
//			{Title: "Title Name2 Something"},
//			{Title: "Title Name2 Something Something"},
//		}
//
// => below will remove any articles where the key is title and similarity score is >= 0.6
//
// result := RemoveDuplicates(articles, 0.6, "Title")
func RemoveDuplicates[T any](items []T, threshold float64, fieldName string) []T {
	var uniqueItems []T

	for i, item := range items {
		isDuplicate := false
		for j := 0; j < i; j++ {
			value1 := reflect.ValueOf(item).FieldByName(fieldName).String()
			value2 := reflect.ValueOf(items[j]).FieldByName(fieldName).String()
			similarity := CalculateSimilarity(value1, value2)
			if similarity >= threshold {
				log.Printf("Excluded potential duplicate value \"%s\" and \"%s\" with similarity score %.2f", value1, value2, similarity)
				isDuplicate = true
				break
			}
		}
		if !isDuplicate {
			uniqueItems = append(uniqueItems, item)
		}
	}

	return uniqueItems
}
