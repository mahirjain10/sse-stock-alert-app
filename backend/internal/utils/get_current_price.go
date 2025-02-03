package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/mahirjain_10/stock-alert-app/backend/internal/types"
)


func GetCurrentStockPriceAndTime(TTM types.Ticker,stockData types.StockData) (float64,string,error) {
	// Fetch stock data from external API
	res, err := http.Get( os.Getenv("STOCK_API_URL")+ TTM.TickerToMonitor + "?range=1d&interval=1m")
	fmt.Println(err)
	fmt.Println(res.Body)
	if err != nil {
		// helpers.SendResponse(c, http.StatusInternalServerError, "Failed to fetch stock price,try again", nil, nil, false)
		return 0,"",fmt.Errorf("failed to fetch stock price, try again")
	}
	defer res.Body.Close() // Ensure the response body is closed after reading

	// Decode the JSON response into stockData struct
	if err := json.NewDecoder(res.Body).Decode(&stockData); err != nil {
		fmt.Println(err)
		// c.JSON(http.StatusBadRequest, gin.H{"statusCode": http.StatusBadRequest, "message": "failed to decode response json", "error": err.Error()})
		return 0,"",fmt.Errorf("failed to decode response json")
	}

	// Check if we have valid data
	if len(stockData.Chart.Result) == 0 || len(stockData.Chart.Result[0].Indicators.Quote) == 0 || len(stockData.Chart.Result[0].Indicators.Quote[0].Close) == 0 {
		// c.JSON(http.StatusNotFound, gin.H{"statusCode": http.StatusNotFound, "message": "no data found", "error": nil})
		return 0,"",fmt.Errorf("no data found")
	}

	// Get the latest price and current time
	latestPrice := stockData.Chart.Result[0].Indicators.Quote[0].Close[len(stockData.Chart.Result[0].Indicators.Quote[0].Close)-1]
	currentTime := time.Now().Format("02-01-2006 15:04:05")
	return latestPrice,currentTime ,nil
}