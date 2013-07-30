/*
Package wbdata provides a client for using the World Bank Open Data API.

Access different parts of the World Bank Open Data API using the various
services:
         client := wbdata.NewClient(nil)

         // list all countries
         countries, err := client.Countries.GetCountries()


The full World Bank Open Data API is documented at http://data.worldbank.org/developers/api-overview.
*/

package wbdata

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

const (
	defaultBaseURL = "http://api.worldbank.org"
)

// A Client manages communication with the World Bank Open Data API.
type Client struct {
	client *http.Client

	BaseURL *url.URL

	//Services to talk to different APIs
	Countries  *CountryService
	Sources    *SourcesService
	Topics     *TopicsService
	Indicators *IndicatorService
}

func NewClient() *Client {
	baseURL, _ := url.Parse(defaultBaseURL)
	c := &Client{client: http.DefaultClient, BaseURL: baseURL}
	c.Countries = &CountryService{client: c}
	return c
}

// NewRequest creates an API request. A relative URL can be provided in urlStr,
// in which case it is resolved relative to the BaseURL of the Client.
// Relative URLs should always be specified without a preceding slash.  If
// specified, the value pointed to by body is JSON encoded and included as the
// request body.
func (c *Client) NewRequest(method, urlStr string, body interface{}) (*http.Request, error) {
	v := url.Values{}
	v.Set("format", "json")

	rel, err := url.Parse(urlStr)
	if err != nil {
		return nil, err
	}

	u := c.BaseURL.ResolveReference(rel)

	url := fmt.Sprintf("%s?%s", u, v.Encode())
	log.Println(url)

	buf := new(bytes.Buffer)
	if body != nil {
		err := json.NewEncoder(buf).Encode(body)
		if err != nil {
			return nil, err
		}
	}

	req, err := http.NewRequest(method, url, buf)
	if err != nil {
		return nil, err
	}

	return req, nil
}

// Do sends an API request and returns the API response.  The API response is
// decoded and stored in the value pointed to by v, or returned as an error if
// an API error has occurred.
func (c *Client) Do(req *http.Request, v *[]interface{}) (*http.Response, error) {
	resp, err := c.client.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if v != nil {
		//err = json.NewDecoder(resp.Body).Decode(v)
		err = json.Unmarshal(body, v)
	}

	return resp, err
}

type ErrorResponse struct {
	Response *http.Response // HTTP response that caused this error
	Message  string         `json:message` // error message
}

func (r *ErrorResponse) Error() string {
	return fmt.Sprintf("%v %v: %d %+v",
		r.Response.Request.Method, r.Response.Request.URL,
		r.Response.StatusCode, r.Message)
}

func CheckResponse(r *http.Response) error {
	errorResponse := &ErrorResponse{Response: r}
	return errorResponse

}
