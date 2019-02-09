package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheMickeyMike/insta-check/pkg/client"
	"github.com/TheMickeyMike/insta-check/pkg/config"
)

type Instagram struct {
	url    string
	client client.HttpClient
}

func NewInstagram(config *config.InstagramConfig, httpClient client.HttpClient) *Instagram {
	return &Instagram{
		url:    config.URL,
		client: httpClient,
	}
}

func (i *Instagram) UsernameIsAvailable(username string) (bool, error) {
	resp, err := i.client.PostForm(i.url, NewRegistrationForm(username))
	if err != nil {
		return false, fmt.Errorf("Can't post form data, error: %v", err)
	}
	defer resp.Body.Close()

	switch resp.StatusCode {
	case http.StatusOK:
		var registerFormResult RegisterFormResponse
		if err := json.NewDecoder(resp.Body).Decode(&registerFormResult); err != nil {
			return false, fmt.Errorf("Can't unmarshall Instagram response, error: %v", err)
		}
		return registerFormResult.UsernameIsOK(), nil
	case http.StatusTooManyRequests:
		return false, fmt.Errorf("Oh Snap! Code: %d Instagram is rate limiting us", http.StatusTooManyRequests)
	default:
		return false, nil
	}
}
