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
	//"strconv"
	"strings"
	//"github.com/beevik/etree"
)

func main() {
	expediaUrl := "https://www.expedia.com/Flights-Search?trip=oneway&leg1=from:nyc,to:mia,departure:09/01/2019TANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass%3Aeconomy&mode=search&origref=www.expedia.com"
	//doc := etree.NewDocument() kexpediaUrl := "https://www.expedia.com/Flights-Search?trip=oneway&leg1=from:nyc,to:mia,departure:09/01/2019TANYT&passengers=adults:1,children:0,seniors:0,infantinlap:Y&options=cabinclass%3Aeconomy&mode=search&origref=www.expedia.com"
	//fmt.Println(expediaUrl)
	//url := "https://www.example.com"
	//httpClientTest(url)
	expediaFlightSearch(expediaUrl)
}

//func main() {
//	in := []byte(`{ "votes": { "option_A": "3" } }`)
//	var raw map[string]interface{}
//	json.Unmarshal(in, &raw)
//	raw["count"] = 1
//	out, _ := json.Marshal(raw)
//	println(string(out))
//}

func unMarshalData(data string) {
	fmt.Println(data)
}

func expediaFlightSearch(expediaUrl string) {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", expediaUrl, nil)
	req.Header.Set("User-Agent", "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/70.0.3538.77 Safari/537.36")
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
		//log.Printf(bodyString)
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
			//log.Println("Found:", value)
			//s, err := strconv.Unquote(value)
			var content Content
			err := json.Unmarshal([]byte(value), &content)
			fmt.Println(err)
			fmt.Println(content.Content)
		}
	}
}

type Content struct {
	Content string
}


//	in := []byte(`{ "votes": { "option_A": "3" } }`)
//	var raw map[string]interface{}
//	json.Unmarshal(in, &raw)
//	raw["count"] = 1
//	out, _ := json.Marshal(raw)
//	println(string(out))
//}
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