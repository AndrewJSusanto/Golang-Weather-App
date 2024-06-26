package main

// API KEY: da79587af3d99005cb7e617471c648d0
import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
	"text/template"
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

var temp *template.Template
var queryText string

const api_request = "http://api.weatherapi.com/v1/current.json?key=44bddb1c26774e7882f22134241706&q="

func fetchAPIData(url string) (Response, error) {
	response, err := http.Get(url)
	if err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)

	// do something with responseObject that assigns a struct to change visuals for background color after api data is fetched.

	return responseObject, nil
}

func formHandler(w http.ResponseWriter, req *http.Request) {
	// want to target id = queryInput with name = query
	temp = template.Must(template.ParseFiles("index.gohtml"))
	queryText := req.FormValue("query")

	queryText = strings.ReplaceAll(strings.TrimSpace(queryText), " ", "")
	if len(queryText) == 0 {
		log.Println("Empty Query")
	}
	data, err := fetchAPIData(api_request + queryText)
	if err != nil {
		log.Fatal(err)
	}
	temp.Execute(w, data)
	// fmt.Println(responseObject.Location.Region)
	// fmt.Println(responseObject.Location.Country)
	// fmt.Println(responseObject.Location.Time)
	// fmt.Printf("Latitude: %f\t\n", responseObject.Location.Lat)
	// fmt.Printf("Longitude: %f\t\n", responseObject.Location.Long)

	// fmt.Println(responseObject.Current.Conditions.Descript)
}

func main() {
	//reader := bufio.NewReader(os.Stdin)

	http.HandleFunc("/", formHandler)
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("HTTP status 500 - Internal server error: %s", err)
	}

	//querytext, _ = reader.ReadString('\n') // takes Input from user for query text including spaces until newline
	// querytext = strings.ReplaceAll(strings.TrimSpace(querytext), " ", "")
	// fmt.Println(GETREQ + querytext)
}
