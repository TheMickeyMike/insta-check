package service

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/TheMickeyMike/insta-check/pkg/client"
)

const registerAjaxURL = "https://www.instagram.com/accounts/web_create_ajax/attempt/"

type Instagram struct {
	client client.HttpClient
}

func NewInstagram(httpClient client.HttpClient) *Instagram {
	return &Instagram{
		client: httpClient,
	}
}

func (i *Instagram) UsernameIsAvailable(username string) (bool, error) {
	resp, err := i.client.PostForm(registerAjaxURL, NewRegistrationForm(username))
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
