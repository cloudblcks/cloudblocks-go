package control

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
)

// The URL of the API
const API_URL = "https://creds.cloudblocks.dev"

type Client struct {
	username string
	password string
	// Url is the base url for the REST API.
	Url string
}

// NewClient creates a new REST client.
func NewClient(username, password string) (Client, error) {
	return NewClientWithURL(username, password, API_URL)
}

// / NewClientWithURL is the same as NewClient but also takes a custom url.
func NewClientWithURL(username, password, url string) (Client, error) {
	client := Client{
		username: username,
		password: password,
		Url:      url,
	}
	return client, nil
}

func (c *Client) RequestCredentials(originId, targetId string) (string, error) {
	var out string
	err := c.request("POST", "/get_creds", map[string]string{
		"origin_id": originId,
		"target_id": targetId,
	}, &out)
	if err != nil {
		return "", err
	}
	return out, nil
}

func (c *Client) request(method, path string, in, out interface{}) error {
	var inR io.Reader
	if in != nil {
		inData, err := json.Marshal(in)
		if err != nil {
			return err
		}
		inR = bytes.NewReader(inData)
	}
	req, err := http.NewRequest(method, c.Url+path, inR)
	if err != nil {
		return err
	}
	req.SetBasicAuth(c.username, c.password)
	if in != nil {
		req.Header.Set("Content-Type", "application/json")
	}
	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode < 200 || res.StatusCode >= 300 {
		_, err := io.ReadAll(res.Body)
		return err
	}
	if out != nil {
		return json.NewDecoder(res.Body).Decode(out)
	}
	return nil
}
