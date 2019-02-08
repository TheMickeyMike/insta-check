package client_test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/TheMickeyMike/insta-check/pkg/client"
	"github.com/stretchr/testify/assert"
)

const testPath = "/some/path"

func TestIfParseFormPostProperForm(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		assert.Equal(t, req.URL.String(), testPath)
		assert.Equal(t, req.FormValue("field1"), "value1")
		rw.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	testURL := fmt.Sprintf("%s%s", server.URL, testPath)
	form := url.Values{
		"field1": {"value1"},
	}

	trickyHTTPClient := client.NewTrickyHTTP()
	resp, err := trickyHTTPClient.PostForm(testURL, form)
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
