package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/joho/godotenv"
)

var apiKey string
var port string

func main() {
	_ = godotenv.Load()
	apiKey = os.Getenv("alpha_vantage_api_key")
	port = os.Getenv("port")
	http.HandleFunc("/status", status)
	http.HandleFunc("/info", info)
	http.ListenAndServe(":"+port, nil)
}
func status(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("good"))
}
func info(w http.ResponseWriter, r *http.Request) {
	symbol := r.URL.Query().Get("symbol")
	url := fmt.Sprintf(
		"https://www.alphavantage.co/query?function=GLOBAL_QUOTE&symbol=%s&apikey=%s", symbol, apiKey)
	client := &http.Client{Timeout: 10 * time.Second}
	resp, err := client.Get(url)
	if err != nil {
		http.Error(w, "Oshibochka providera", 500)
		return
	}
	defer resp.Body.Close()
	var data map[string]interface{}
	_ = json.NewDecoder(resp.Body).Decode(&data)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}
