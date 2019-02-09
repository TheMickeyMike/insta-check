package service_test

import (
	"github.com/TheMickeyMike/insta-check/pkg/config"
	"bytes"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/TheMickeyMike/insta-check/pkg/client/mocks"
	"github.com/TheMickeyMike/insta-check/pkg/service"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/assert"
)

const testUsername = "some_username"
const responseWithNotValidUsername = `
{
    "errors": {
        "username": [
            {
                "message": "This field is required.",
                "code": "username_required"
            }
        ]
    }
}`
const responseWithValidUsername = `
{
    "errors": {
        "username": []
    }
}`

func TestUsernameIsAvailableReturnsFalseIfInstgramResturnValidationError(t *testing.T) {
	instaConfig := &config.InstagramConfig{URL: "http://example.com/some/path"}
	form := service.NewRegistrationForm(testUsername)

	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(responseWithNotValidUsername))),
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	httpClientMock := mocks.NewMockHttpClient(mockCtrl)
	httpClientMock.EXPECT().PostForm(instaConfig.URL, form).Return(response, nil)

	instagram := service.NewInstagram(instaConfig,httpClientMock)
	result, err := instagram.UsernameIsAvailable(testUsername)
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestUsernameIsAvailableReturnsTrue(t *testing.T) {
	instaConfig := &config.InstagramConfig{URL: "http://example.com/some/path"}
	form := service.NewRegistrationForm(testUsername)

	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(responseWithValidUsername))),
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	httpClientMock := mocks.NewMockHttpClient(mockCtrl)
	httpClientMock.EXPECT().PostForm(instaConfig.URL, form).Return(response, nil)

	instagram := service.NewInstagram(instaConfig,httpClientMock)
	result, err := instagram.UsernameIsAvailable(testUsername)
	assert.NoError(t, err)
	assert.True(t, result)
}
