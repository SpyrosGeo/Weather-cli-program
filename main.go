package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Response struct {
	GenerationtimeMs float64 `json:"generationtime_ms"`
	HourlyUnits      struct {
		Time          string `json:"time"`
		Temperature2M string `json:"temperature_2m"`
	} `json:"hourly_units"`
	Hourly struct {
		Time          []string  `json:"time"`
		Temperature2M []float64 `json:"temperature_2m"`
	} `json:"hourly"`
	Elevation        float64 `json:"elevation"`
	Latitude         int     `json:"latitude"`
	Longitude        float64 `json:"longitude"`
	UtcOffsetSeconds int     `json:"utc_offset_seconds"`
}

func main() {
	// dt := time.Now()
	// currentTime := dt.String()
	// log.Printf(currentTime)
	resp, err := http.Get("https://api.open-meteo.com/v1/forecast?latitude=37.9792&longitude=23.7166&hourly=temperature_2m")

	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	var result Response
	if err := json.Unmarshal(body, &result); err != nil {
		fmt.Println("Can not unmarshal JSON")
	}
	for _, rec := range result.Hourly {
		return
		// fmt.Println(rec.)
	}
}
