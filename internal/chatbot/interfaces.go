package chatbot;

// Chatbot interface defines methods for chatbot functionalities.
interface Chatbot {
    void Chat(String message);
    void SetContext(String context);
}

// FinancialChatbot interface extends Chatbot with financial analysis methods.
interface FinancialChatbot extends Chatbot {
    void AnalyzeFinancials(String data);
    void FetchFinancialData(String source);
    double PredictStockPrice(String stockSymbol);
}