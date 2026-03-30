package finance

// StockAnalyzer interface for analyzing stocks.
type StockAnalyzer interface {
    Analyze(stock string) (float64, error)
}

// StockPredictor interface for predicting stock prices.
type StockPredictor interface {
    Predict(stock string) (float64, error)
}

// PortfolioAnalyzer interface for analyzing a stock portfolio.
type PortfolioAnalyzer interface {
    AnalyzePortfolio(stocks []string) (float64, error)
}