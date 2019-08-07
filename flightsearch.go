package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"golang.org/x/net/html"
	"gopkg.in/xmlpath.v2"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)


type FlightData struct {

	DepartureCity string
	DepartureCode string
	DepartureTimeIsoStr string
	ArrivalCity string
	ArrivalCode string
	ArrivalTimeIsoStr string
	Price float64
	Airline string

}

type Content struct {
	Content string
}

type SearchCriteria struct {
	DepartureCode string
	ArrivalCode string
	DepartureDate string
}


func readContent(content string) {
	var raw map[string]interface{}
	json.Unmarshal([]byte(content), &raw)
	legs := raw["legs"].(map[string]interface{})
	for _, v := range legs {

		legData := v.(map[string]interface{})
		departureCity := legData["departureLocation"].(map[string]interface{})["airportCity"].(string)
		departureCode := legData["departureLocation"].(map[string]interface{})["airportCode"].(string)
		departureTimeIsoStr := legData["departureTime"].(map[string]interface{})["isoStr"].(string)
		arrivalCity := legData["arrivalLocation"].(map[string]interface{})["airportCity"].(string)
		arrivalCode := legData["arrivalLocation"].(map[string]interface{})["airportCode"].(string)
		arrivalTimeIsoStr := legData["arrivalTime"].(map[string]interface{})["isoStr"].(string)
		price := legData["price"].(map[string]interface{})["totalPriceAsDecimal"].(float64)
		airline := legData["carrierSummary"].(map[string]interface{})["airlineName"].(string)


		flightData := FlightData{
			DepartureCity: departureCity,
			DepartureCode: departureCode,
			DepartureTimeIsoStr: departureTimeIsoStr,
			ArrivalCity: arrivalCity,
			ArrivalCode: arrivalCode,
			ArrivalTimeIsoStr: arrivalTimeIsoStr,
			Price: price,
			Airline: airline,
		}

		flightDataJson, _ := json.Marshal(flightData)
		fmt.Println(string(flightDataJson))

	}
}

func expediaFlightSearch(expediaUrl string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", expediaUrl, nil)
	android74Ua := "Mozilla/5.0 (Linux; Android 9; ONEPLUS A5010) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/74.0.3729.157 Mobile Safari/537.36 Brave/74"
	req.Header.Set("User-Agent", android74Ua)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode == http.StatusOK {
		bodyBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatal(err)
		}
		bodyString := string(bodyBytes)
		reader := strings.NewReader(bodyString)
		root, err := html.Parse(reader)
		if err != nil {
			log.Fatal(err)
		}
		var b bytes.Buffer
		html.Render(&b, root)
		fixedHtml := b.String()

		reader = strings.NewReader(fixedHtml)
		xmlroot, xmlerr := xmlpath.ParseHTML(reader)

		if xmlerr != nil {
			log.Fatal(xmlerr)
		}

		var xpath string
		xpath = "//script[@id='cachedResultsJson']//text()"
		path := xmlpath.MustCompile(xpath)
		if value, ok := path.String(xmlroot); ok {
			var content Content
			err := json.Unmarshal([]byte(value), &content)
			if err != nil {
				log.Fatal(err)
			}
			readContent(content.Content)
		}
	}
}


func main() {

	search := SearchCriteria{
		DepartureCode: "nyc",
		ArrivalCode: "sfo",
		DepartureDate: "11/01/2019",
	}

	expediaUrl := fmt.Sprintf(
		"https://www.expedia.com/Flights-Search?trip=oneway&leg1=from:%s,to:%s,departure:%sTANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass%3Aeconomy&mode=search&origref=www.expedia.com",
		search.DepartureCode,
		search.ArrivalCode,
		search.DepartureDate)
	expediaFlightSearch(expediaUrl)
}

