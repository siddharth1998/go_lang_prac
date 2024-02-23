package myPkg

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type HostLocation struct {
	City string `json:"city"`
	RegionCode string `json:"region_code"`
	AreaCode int `json:"area_code"`
	Longitude float32 `json:"longitude"`
	CountryCode3 string `json:"country_code3"`
	CountryName string `json:"country_name"`
	PostalCode string `json:"postal_code"`
	DMACode int `json:"dma_code"`
	CountryCode string `json:"country_code"`
	Latitude float32 `json:"latitude"`
}

type Host struct{
	OS string `json:"os"`
	Timestamp string `json:"timestamp"`
	ISP string `json:"isp"`
	ASN string `json:"asn"`
	Hostnames []string `json:"hostnames"`
	Location HostLocation `json:"location"`
	IP int64 `json:"ip"`
	Domains []string `json:"domains"`
	Org string `json:"org"`
	Data string `json:"data"`
	Port int `json:"port"`
	IPString string `json:"ip_str"`
}

type HostSearch struct{
	Matches []Host `json:"matches"`
}

func (s *Client) HostSearch(q string)(*HostSearch,error){
	stringy:=fmt.Sprintf("%sshodan/host/search?key=%s&query=%s",BaseURL,s.apiKey,q)
	res,err := http.Get(stringy)
	if err!=nil{
		return nil,err
	}
	defer res.Body.Close()
	var returnValue HostSearch
	if err:= json.NewDecoder(res.Body).Decode(&returnValue);err!=nil{
		log.Fatalln("Can encode and decode")
		return nil,err
	}
	return &returnValue,nil
} 

