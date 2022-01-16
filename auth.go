package variti

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

// SignIn obtain a new token from user and password, return AuthResponse and error
func (c *Client) SignIn() (*AuthResponse, error) {
	if c.Auth.Username == "" || c.Auth.Password == "" {
		return nil, fmt.Errorf("SignIn: define username and password")
	}

	rb, err := json.Marshal(c.Auth)
	if err != nil {
		return nil, fmt.Errorf("SignIn: json marshal failed: %w", err)
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/login", c.HostURL), strings.NewReader(string(rb)))
	if err != nil {
		return nil, fmt.Errorf("SignIn: login request failed: %w", err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return nil, fmt.Errorf("SignIn: login request failed: %w", err)
	}

	ar := AuthResponse{}
	err = json.Unmarshal(body, &ar)
	if err != nil {
		return nil, fmt.Errorf("SignIn: json unmarshal failed: %w", err)
	}

	return &ar, nil
}

// SignOut revoke the token for a user (future use)
func (c *Client) SignOut() error {
	req, err := http.NewRequest("POST", fmt.Sprintf("%s/auth/logout", c.HostURL), strings.NewReader(string("")))
	if err != nil {
		return fmt.Errorf("SignOut: logout failed: %w", err)
	}

	body, err := c.doRequest(req)
	if err != nil {
		return fmt.Errorf("SignOut: logout request failed: %w", err)
	}

	if string(body) != "Signed out user" {
		return errors.New(string(body))
	}

	return nil
}
