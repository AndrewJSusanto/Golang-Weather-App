package main

// API KEY: da79587af3d99005cb7e617471c648d0
import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

type Response struct {
	Location Location `json:"location"`
	Current  Current  `json:"current"`
}

type Location struct {
	Name    string  `json:"name"`
	Region  string  `json:"region"`
	Country string  `json:"country"`
	Lat     float32 `json:"lat"`
	Long    float32 `json:"lon"`
	Time    string  `json:"localtime"`
}

type Current struct {
	Updated    string    `json:"last_updated"`
	Conditions Condition `json:"condition"`
	IsDay      bool      `json:"is_day"`

	TempC      float32 `json:"temp_c"`
	TempF      float32 `json:"temp_f"`
	FeelsLikeC float32 `json:"feelslike_c"`
	FeelsLikeF float32 `json:"feelslike_f"`

	WindMPH    float32 `json:"wind_mph"`
	WindKPH    float32 `json:"wind_kph"`
	WindDegree int     `json:"wind_degree"`
	WindDir    string  `json:"wind_dir"`
	WindChillC float32 `json:"windchill_c"`
	WindChillF float32 `json:"windchill_f"`

	PressureMB float32 `json:"pressure_mb"`
	PressureIN float32 `json:"pressure_in"`
	VisKM      float32 `json:"vis_km"`
	VisMI      float32 `json:"vis_miles"`

	PrecipMM float32 `json:"precip_mm"`
	PrecipIN float32 `json:"precip_in"`
	Humidity int     `json:"humidity"`
}

type Condition struct {
	Descript string `json:"text"`
	IconURL  string `json:"icon"`
	Code     int    `json:"code"`
}

func main() {
	const GETREQ = "http://api.weatherapi.com/v1/current.json?key=44bddb1c26774e7882f22134241706&q="
	var querytext string
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Please input your query (latitude, longitude; city name)")

	querytext, _ = reader.ReadString('\n') // takes Input from user for query text including spaces until newline
	querytext = strings.ReplaceAll(strings.TrimSpace(querytext), " ", "")
	// fmt.Println(GETREQ + querytext)
	response, err := http.Get(GETREQ + querytext)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	fmt.Println(responseObject.Location.Region)
	fmt.Println(responseObject.Location.Country)
	fmt.Println(responseObject.Location.Time)
	fmt.Printf("Latitude: %f\t\n", responseObject.Location.Lat)
	fmt.Printf("Longitude: %f\t\n", responseObject.Location.Long)

	fmt.Println(responseObject.Current.Conditions.Descript)

}
