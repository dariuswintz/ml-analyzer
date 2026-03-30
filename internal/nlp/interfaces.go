// Package nlp defines the interfaces for various NLP components.

package nlp

// SentimentAnalyzer defines the methods for analyzing sentiment.
type SentimentAnalyzer interface {
    Analyze(text string) (float64, error)
}

// TextSummarizer defines the methods for summarizing text.
type TextSummarizer interface {
    Summarize(text string, maxLength int) (string, error)
}

// TextGenerator defines the methods for generating text.
type TextGenerator interface {
    Generate(prompt string, maxLength int) (string, error)
}