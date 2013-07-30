package wbdata

// CountryService provides access to the Indicator related queries
// in the World Bank Open Data API.
//
// World Bank Open Data API docs: http://data.worldbank.org/node/203
type IndicatorService struct {
	client *Client
}
