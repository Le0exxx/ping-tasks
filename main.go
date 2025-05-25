package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"sort"
	"strconv"
	"Le0exxx/ping-tasks/pkg/utils"
)

type TimeSeriesResponse struct {
	TimeSeries map[string]map[string]string `json:"Time Series (Daily)"`
}

type Response struct {
	Symbol        string             `json:"symbol"`
	DaysRequested int                `json:"days_requested"`
	Data          []DayData          `json:"data"`
	AverageClose  float64            `json:"average_close"`
}

type DayData struct {
	Date  string  `json:"date"`
	Close float64 `json:"close"`
}

func main() {
	http.HandleFunc("/rabbit", handlePrices)

	port := "8080"
	fmt.Println("Server running on port", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func handlePrices(w http.ResponseWriter, r *http.Request) {
	symbol := os.Getenv("SYMBOL")
	ndaysStr := os.Getenv("NDAYS")
	apiKey := os.Getenv("APIKEY")

	if symbol == "" || ndaysStr == "" || apiKey == "" {
		utils.JsonError(w, http.StatusBadRequest, "SYMBOL, NDAYS and APIKEY environment variables must be set")
		return
	}

	ndays, err := strconv.Atoi(ndaysStr)
	if err != nil || ndays <= 0 {
		utils.JsonError(w, http.StatusBadRequest, "NDAYS must be a positive integer")
		return
	}

	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=TIME_SERIES_DAILY&symbol=%s&apikey=%s",
		symbol, apiKey,
	)

	resp, err := http.Get(url)
	if err != nil {
		utils.JsonError(w, http.StatusInternalServerError, "Failed to fetch stock data")
		return
	}
	defer resp.Body.Close()

	var tsr TimeSeriesResponse
	if err := json.NewDecoder(resp.Body).Decode(&tsr); err != nil {
		http.Error(w, "Failed to parse API response", http.StatusInternalServerError)
		return
	}

	// Extract and sort the dates (descending)
	dates := make([]string, 0, len(tsr.TimeSeries))
	for date := range tsr.TimeSeries {
		dates = append(dates, date)
	}
	sort.Sort(sort.Reverse(sort.StringSlice(dates)))

	// Get the last NDAYS
	var data []DayData
	var sum float64
	count := 0
	for _, date := range dates {
		if count >= ndays {
			break
		}
		day := tsr.TimeSeries[date]
		closeStr, ok := day["4. close"]
		if !ok {
			log.Printf("Missing close value for date %s", date)
			continue
		}
		closeVal, err := strconv.ParseFloat(closeStr, 64)
		if err != nil {
			log.Printf("Invalid close value for date %s: %v", date, err)
			continue
		}
		data = append(data, DayData{Date: date, Close: closeVal})
		sum += closeVal
		count++
	}

	result := Response{
		Symbol:        symbol,
		DaysRequested: ndays,
		Data:          data,
		AverageClose:  sum / float64(len(data)),
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(result)
}
