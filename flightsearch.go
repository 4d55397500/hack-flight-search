package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	//"github.com/beevik/etree"
)

func main() {

	fmt.Printf("foo")
	//doc := etree.NewDocument()

//	url := "https://www.expedia.com/Flights-Search?trip=oneway&leg1=from:nyc,to:mia,departure:04/01/2017TANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass%3Aeconomy&mode=search&origref=www.expedia.com"
	url := "https://www.example.com"
	httpClientTest(url)
}

func expediaFlightSearch(url string) {

}

func httpClientTest(url string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", url, nil)
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
		log.Printf(bodyString)
	}
}