package finance

import (
	"fmt"
)

// SimpleStockAnalyzer provides methods to analyze stock data.
type SimpleStockAnalyzer struct {
	StockData []float64
}

// GetStockData fetches or simulates stock data.
func (s *SimpleStockAnalyzer) GetStockData() {
	// Simulate fetching stock data
	s.StockData = []float64{150.0, 153.5, 149.0, 155.0, 158.0, 156.5}
}

// CalculateMovingAverage calculates the moving average of the stock data for a given period.
func (s *SimpleStockAnalyzer) CalculateMovingAverage(period int) float64 {
	if len(s.StockData) < period {
		fmt.Println("Not enough data")
		return 0.0
	}
	
	sum := 0.0
	for _, price := range s.StockData[len(s.StockData)-period:] {
		sum += price
	}
	return sum / float64(period)
}

// AnalyzeVolatility analyzes the volatility of the stock data.
func (s *SimpleStockAnalyzer) AnalyzeVolatility() float64 {
	mean := 0.0
	
	// Calculate mean price
	for _, price := range s.StockData {
		mean += price
	}
	mean /= float64(len(s.StockData))

	varianceSum := 0.0
	
	// Calculate variance
	for _, price := range s.StockData {
		varianceSum += (price - mean) * (price - mean)
	}
	variance := varianceSum / float64(len(s.StockData))

	return variance
}