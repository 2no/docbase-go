package docbase

import (
	"encoding/json"
	"fmt"
)

type Group struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

func (c Client) Groups() ([]Group, *RateLimit, error) {
	path := fmt.Sprintf("/teams/%s/groups", c.Team)
	url, err := c.assembleURL(path)
	if err != nil {
		return nil, nil, err
	}

	body, rateLimit, err := c.get(url)
	if err != nil {
		return nil, rateLimit, err
	}

	var groups []Group
	json.Unmarshal(body, &groups)
	return groups, rateLimit, nil
}
