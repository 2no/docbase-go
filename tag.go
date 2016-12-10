package docbase

import (
	"encoding/json"
	"fmt"
)

type Tag struct {
	Name string `json:"name"`
}

func (c Client) Tags() ([]Tag, *RateLimit, error) {
	path := fmt.Sprintf("/teams/%s/tags", c.Team)
	url, err := c.assembleURL(path)
	if err != nil {
		return nil, nil, err
	}

	body, rateLimit, err := c.get(url)
	if err != nil {
		return nil, rateLimit, err
	}

	var tags []Tag
	json.Unmarshal(body, &tags)
	return tags, rateLimit, nil
}
