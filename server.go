package main

import (
	"encoding/csv"
	"encoding/json"
	"log"
	"net/http"
	"os"
)

// Date, Close/Last, Volume, Open, High, Low
// 02/28/2020, $273.36, 106721200, $257.26, $278.41, $256.37


type Stock struct {
	Date   string
	Close  float64
	Volume int
	Open   float64
	High   float64
	Low    float64
}

func handler(w http.ResponseWriter, r *http.Request) {
	f, err := os.Open("prices.csv")
	if err != nil {
		log.Fatal(err)
	}
	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	prices, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(prices)

}

func main() {

	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
