package wbdata

import (
	"fmt"
)

type CountryService struct {
	client *Client
}

type Country struct {
	Name                string
	CapitalCity         string
	Iso2Code            string
	Longitude, Latitude string
}

type CountryHeader struct {
	Page    int
	Pages   int
	PerPage string
	Total   int
}

type CountryListOptions struct {
	Page    int
	PerPage int
}

// GetSources retrives Catalog Data Sources from the World Bank API
// It return s a slice of Source
func (c *CountryService) GetCountries() ([]Country, error) {
	countryHeader := CountryHeader{}
	country := []Country{}

	req, err := c.client.NewRequest("GET", "countries", nil)
	if err != nil {
		return nil, err
	}

	_, err = c.client.Do(req, &[]interface{}{&countryHeader, &country})

	return country, err
}

func (c *CountryService) GetCountry(countryId string) ([]Country, error) {
	countryHeader := CountryHeader{}
	country := []Country{}

	u := fmt.Sprintf("countries/%v", countryId)
	req, err := c.client.NewRequest("GET", u, nil)

	if err != nil {
		return nil, err
	}
	_, err = c.client.Do(req, &[]interface{}{&countryHeader, &country})

	return country, err

}
