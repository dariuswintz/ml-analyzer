package nlp

import (
	"sort"
)

// SimpleTextSummarizer provides methods for extractive text summarization.
type SimpleTextSummarizer struct {
	texts []string
}

// NewSimpleTextSummarizer initializes a new SimpleTextSummarizer.
func NewSimpleTextSummarizer(texts []string) *SimpleTextSummarizer {
	return &SimpleTextSummarizer{texts: texts}
}

// Summarize returns the top n sentences based on their importance.
func (s *SimpleTextSummarizer) Summarize(n int) []string {
	// In a real implementation, you would compute sentence weights based on importance.
	// For this example, we simply return the first n sentences.
	if n > len(s.texts) {
		n = len(s.texts)
	}
	return s.texts[:n]
}

// Sort sentences by their importance (this is a mock sorting function).
func (s *SimpleTextSummarizer) SortSentencesByImportance(importance []float64) []string {
	sortedTexts := make([]string, len(s.texts))
	copy(sortedTexts, s.texts)
	// Sort based on importance - just a placeholder example
	sort.Slice(sortedTexts, func(i, j int) bool {
		return importance[i] > importance[j]
	})
	return sortedTexts
}
