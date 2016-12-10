package docbase

import (
	"errors"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"path"
)

const Version = "0.0.1"

const (
	defaultURL = "https://api.docbase.io"
	userAgent  = "DocBase GoLang " + Version
)

type Client struct {
	Team        string
	accessToken string
	url         string
}

type ClientValues struct {
	AccessToken string
	Url         string
	Team        string
}

type RateLimit struct {
	Limit     string
	Remaining string
	Reset     string
}

func NewClient(values *ClientValues) *Client {
	client := &Client{
		Team:        values.Team,
		accessToken: values.AccessToken,
		url:         values.Url,
	}
	if len(client.url) == 0 {
		client.url = defaultURL
	}
	return client
}

func (c Client) assembleURL(pathStr string) (string, error) {
	u, err := url.Parse(c.url)
	if err != nil {
		return "", err
	}

	u.Path = path.Join(u.Path, pathStr)
	return u.String(), nil
}

func (c Client) newRequest(method string, urlStr string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, urlStr, body)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Accept", "application/json")
	req.Header.Set("X-DocBaseToken", c.accessToken)
	req.Header.Set("UserAgent", userAgent)
	return req, err
}

func (c Client) request(method string, urlStr string, body io.Reader) ([]byte, *RateLimit, error) {
	req, err := c.newRequest(method, urlStr, body)
	if err != nil {
		return nil, nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, nil, err
	}
	defer resp.Body.Close()

	rateLimit := &RateLimit{
		Limit:     resp.Header.Get("X-RateLimit-Limit"),
		Remaining: resp.Header.Get("X-RateLimit-Remaining"),
		Reset:     resp.Header.Get("X-RateLimit-Reset"),
	}

	if resp.StatusCode != http.StatusOK &&
		resp.StatusCode != http.StatusCreated &&
		resp.StatusCode != http.StatusNoContent {
		return nil, rateLimit, errors.New(http.StatusText(resp.StatusCode))
	}

	byteArray, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, rateLimit, err
	}

	return byteArray, rateLimit, nil
}

func (c Client) get(urlStr string) ([]byte, *RateLimit, error) {
	return c.request("GET", urlStr, nil)
}

func (c Client) post(urlStr string, body io.Reader) ([]byte, *RateLimit, error) {
	return c.request("POST", urlStr, body)
}

func (c Client) delete(urlStr string) (*RateLimit, error) {
	_, rateLimit, err := c.request("DELETE", urlStr, nil)
	return rateLimit, err
}

func (c Client) patch(urlStr string, body io.Reader) ([]byte, *RateLimit, error) {
	return c.request("PATCH", urlStr, body)
}
