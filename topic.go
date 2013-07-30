package wbdata

type TopicsService struct {
	client *Client
}

type TopicHeader struct {
	Page    int
	Pages   int
	PerPage int
	Total   int
}

type Topic struct {
	Id, Value, SourceNote string
}

func (t *TopicsService) GetTopics() ([]Topic, error) {

	header := TopicHeader{}
	topic := []Topic{}

	req, err := t.client.NewRequest("GET", "topics", nil)
	if err != nil {
		return nil, err
	}

	_, err = t.client.Do(req, &[]interface{}{&header, &topic})

	return topic, err

}
