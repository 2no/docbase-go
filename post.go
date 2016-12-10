package docbase

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/url"
)

type Post struct {
	Id        int     `json:"id"`
	Title     string  `json:"title"`
	Body      string  `json:"body"`
	Draft     bool    `json:"draft"`
	Url       string  `json:"url"`
	CreatedAt string  `json:"created_at"`
	Tags      []Tag   `json:"tags"`
	Scope     string  `json:"scope"`
	Groups    []Group `json:"groups"`
	User      User    `json:"user"`
}

type PostValues struct {
	Title  string   `json:"title"`
	Body   string   `json:"body"`
	Notice bool     `json:"notice"`
	Draft  bool     `json:"draft"`
	Tags   []string `json:"tags,omitempty"`
	Scope  string   `json:"scope"`
	Groups []int    `json:"groups,omitempty"`
}

type PostSearchMeta struct {
	PreviousPage string `json:"previous_page"`
	NextPage     string `json:"next_page"`
	Total        int    `json:"total"`
}

type PostSearchResult struct {
	Posts []Post         `json:"posts"`
	Meta  PostSearchMeta `json:"meta"`
}

func NewPostValues() *PostValues {
	return &PostValues{
		Notice: true,
		Scope:  "everyone",
	}
}

func (c Client) CreatePost(values *PostValues) (*Post, *RateLimit, error) {
	path := fmt.Sprintf("/teams/%s/posts", c.Team)
	url, err := c.assembleURL(path)
	if err != nil {
		return nil, nil, err
	}

	byteArray, err := json.Marshal(values)
	if err != nil {
		return nil, nil, err
	}

	body, rateLimit, err := c.post(url, bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, rateLimit, err
	}

	var post *Post
	json.Unmarshal(body, &post)
	return post, rateLimit, nil
}

func (c Client) UpdatePost(id int, values *PostValues) (*Post, *RateLimit, error) {
	path := fmt.Sprintf("/teams/%s/posts/%d", c.Team, id)
	url, err := c.assembleURL(path)
	if err != nil {
		return nil, nil, err
	}

	byteArray, err := json.Marshal(values)
	if err != nil {
		return nil, nil, err
	}

	body, rateLimit, err := c.patch(url, bytes.NewBuffer(byteArray))
	if err != nil {
		return nil, rateLimit, err
	}

	var post *Post
	json.Unmarshal(body, &post)
	return post, rateLimit, nil
}

func (c Client) DeletePost(id int) (*RateLimit, error) {
	path := fmt.Sprintf("/teams/%s/posts/%d", c.Team, id)
	url, err := c.assembleURL(path)
	if err != nil {
		return nil, err
	}

	return c.delete(url)
}

func (c Client) SearchPost(values *url.Values) (*PostSearchResult, *RateLimit, error) {
	path := fmt.Sprintf("/teams/%s/posts", c.Team)
	url, err := c.assembleURL(path)
	if err != nil {
		return nil, nil, err
	}

	body, rateLimit, err := c.get(url + "?" + values.Encode())
	if err != nil {
		return nil, rateLimit, err
	}

	var result *PostSearchResult
	json.Unmarshal(body, &result)
	return result, rateLimit, nil
}
