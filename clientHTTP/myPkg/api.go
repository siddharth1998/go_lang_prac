package myPkg

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type APIInfo struct {
	QueryCredits int    `json:"query_credits"`
	ScanCredits  int    `json:"scan_credits"`
	Telnet       bool   `json:"telnet"`
	Plan         string `json:"plan"`
	HTTPS        bool   `json:"https"`
	Unlocked     bool   `json:"unlocked"`
}

func (s *Client) APIInfo()(*APIInfo,error) {
	// request shodan for the API key information

	var url = fmt.Sprintf("%sapi-info?key=%s", BaseURL, s.apiKey)
	fmt.Println("the url is "+url)
	res, err := http.Get(url)
	defer res.Body.Close()

	if err != nil {
		fmt.Println("Not able to get the api-key information")
		return nil,err
	}

	var apiInfoVal APIInfo
	// Decode the data
	if err := json.NewDecoder(res.Body).Decode(&apiInfoVal); err != nil {
		fmt.Println("Problem in encoding and decoding")
		return nil,err
	}
	return &apiInfoVal,nil
}
