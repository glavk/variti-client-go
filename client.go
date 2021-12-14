package variti

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

// default API URL
const HostURL string = "https://cp.variti.io/api2"

type Client struct {
	HostURL    string
	HTTPClient *http.Client
	Token      string
	Auth       AuthStruct
}

type AuthStruct struct {
	Username string `json:"login"`
	Password string `json:"password"`
}

type AuthResponse struct {
	Message string `json:"msg"`
	Token   string `json:"access_token"`
	ReToken string `json:"refresh_token,omitempty" `
}

// NewClient creates authenticated by username and password client
func NewClient(host, username, password *string) (*Client, error) {
	c := Client{
		HTTPClient: &http.Client{Timeout: 10 * time.Second},
		HostURL:    HostURL,
		Auth: AuthStruct{
			Username: *username,
			Password: *password,
		},
	}

	if host != nil {
		c.HostURL = *host
	}

	ar, err := c.SignIn()
	if err != nil {
		return nil, err
	}

	c.Token = ar.Token

	return &c, nil
}

// doRequest make request req with bearer token, return responce body and error
func (c *Client) doRequest(req *http.Request) ([]byte, error) {
	req.Header.Set("Content-Type", "application/json")

	if c.Token != "" {
		req.Header.Set("Authorization", "Bearer "+c.Token)
	}

	res, err := c.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status: %d, body: %s", res.StatusCode, body)
	}

	return body, err
}
