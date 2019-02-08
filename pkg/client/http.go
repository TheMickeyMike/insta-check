package client

import (
	"net/http"
	"net/url"
	"strings"
)

var xCsrfToken = http.CanonicalHeaderKey("x-csrftoken")
var contentType = http.CanonicalHeaderKey("content-type")
var userAgent = http.CanonicalHeaderKey("user-agent")
var accept = http.CanonicalHeaderKey("accept")
var referer = http.CanonicalHeaderKey("referer")

type HttpClient interface {
	PostForm(url string, data url.Values) (resp *http.Response, err error)
}

type trickyHTTP struct {
	client *http.Client
}

func NewTrickyHTTP() *trickyHTTP {
	return &trickyHTTP{
		client: &http.Client{},
	}
}

func (thc *trickyHTTP) PostForm(url string, data url.Values) (resp *http.Response, err error) {
	req, err := http.NewRequest(http.MethodPost, url, strings.NewReader(data.Encode()))
	if err != nil {
		return nil, err
	}
	req.Header.Add(xCsrfToken, "QqIADcbXeknPT0KUu7GN7n0IJDq9A7oV")
	req.Header.Add(accept, "application/json; charset=utf-8")
	req.Header.Set(contentType, "application/x-www-form-urlencoded")
	req.Header.Set(userAgent, "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_14_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/72.0.3626.96 Safari/537.36")
	req.Header.Set(referer, "https://www.instagram.com/accounts/emailsignup/")
	return thc.client.Do(req)
}
