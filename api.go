package wbdata

import (
	"encoding/xml"
	"io/ioutil"
	"log"
	"net/http"
)

const (
	baseUrl      = "http://api.worldbank.org"
	sourcesUrl   = baseUrl + "/sources"
	countriesUrl = baseUrl + "/countries"
)

type Wbdata struct{}

type Source struct {
	Name        string `xml:"name"`
	Description string `xml:"description"`
	Url         string `xml:"url"`
}

type Sources struct {
	XMLName    xml.Name `xml"sources"`
	SourceList []Source `xml:"source"`
}

type Country struct {
	Name        string `xml:"name"`
	CapitalCity string `xml:"capitalCity"`
}

type Countries struct {
	XMLName     xml.Name  `xml"countries"`
	CountryList []Country `xml:"country"`
}

func GetSources() (*Sources, error) {
	resp, err := http.Get(sourcesUrl)
	if err != nil {
		log.Fatalf("error %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	s := new(Sources)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	//log.Printf("body %v", string(body))

	xml.Unmarshal(body, &s)
	return s, nil

}

func GetCountries() (*Countries, error) {
	resp, err := http.Get(countriesUrl)
	if err != nil {
		log.Fatalf("error %v", err)
		return nil, err
	}

	defer resp.Body.Close()
	c := new(Countries)
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	//log.Printf("body %v", string(body))

	xml.Unmarshal(body, &c)
	return c, nil
}
