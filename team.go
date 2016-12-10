package docbase

import "encoding/json"

type Team struct {
	Domain string `json:"domain"`
	Name   string `json:"name"`
}

func (c Client) Teams() ([]Team, *RateLimit, error) {
	url, err := c.assembleURL("/teams")
	if err != nil {
		return nil, nil, err
	}

	body, rateLimit, err := c.get(url)
	if err != nil {
		return nil, rateLimit, err
	}

	var teams []Team
	json.Unmarshal(body, &teams)
	return teams, rateLimit, nil
}
