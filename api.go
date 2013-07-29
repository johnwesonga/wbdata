package wbdata

import (
	"encoding/json"
	"log"
)

const (
	baseUrl      = "http://api.worldbank.org"
	sourcesUrl   = baseUrl + "/sources"
	countriesUrl = baseUrl + "/countries"
)

type DataSources struct {
	Page    string
	Pages   string
	PerPage string
	Total   string
}

type Source struct {
	Name string
}

type Country struct {
	Name        string
	CapitalCity string
}

type CountriesSource struct {
	Page    int
	Pages   int
	PerPage string
	Total   int
}

func GetSources() []Source {
	ds := DataSources{}
	s := []Source{}

	body, err := FetchUrl(sourcesUrl)

	if err != nil {
		log.Fatalf("error fetching url %v", err)
	}

	if err := json.Unmarshal(body, &[]interface{}{&ds, &s}); err != nil {
		log.Fatalf("error %v", err)
	}

	return s

}

func GetCountries() []Country {
	cs := CountriesSource{}
	cn := []Country{}

	body, err := FetchUrl(countriesUrl)

	if err != nil {
		log.Fatalf("error fetching url %v", err)
	}

	if err := json.Unmarshal(body, &[]interface{}{&cs, &cn}); err != nil {
		log.Fatalf("error %v", err)
	}

	return cn

}
