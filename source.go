package wbdata

type SourcesList struct {
	client *Client
}

type SourceHeader struct {
	Page    string
	Pages   string
	PerPage string
	Total   string
}

type Source struct {
	Id          string
	Name        string
	Description string
	Url         string
}

func (s *SourcesList) GetSources() ([]Source, error) {
	sourceHeader := SourceHeader{}
	source := []Source{}

	req, err := s.client.NewRequest("GET", "sources", nil)
	if err != nil {
		return nil, err
	}

	_, err = s.client.Do(req, &[]interface{}{&sourceHeader, &source})

	return source, err

}
