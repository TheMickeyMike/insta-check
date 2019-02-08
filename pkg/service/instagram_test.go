package service_test

import (
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
const testURL = "https://www.instagram.com/accounts/web_create_ajax/attempt/"
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
	form := service.NewRegistrationForm(testUsername)

	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(responseWithNotValidUsername))),
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	httpClientMock := mocks.NewMockHttpClient(mockCtrl)
	httpClientMock.EXPECT().PostForm(testURL, form).Return(response, nil)

	instagram := service.NewInstagram(httpClientMock)
	result, err := instagram.UsernameIsAvailable(testUsername)
	assert.NoError(t, err)
	assert.False(t, result)
}

func TestUsernameIsAvailableReturnsTrue(t *testing.T) {
	form := service.NewRegistrationForm(testUsername)

	response := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(responseWithValidUsername))),
	}

	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	httpClientMock := mocks.NewMockHttpClient(mockCtrl)
	httpClientMock.EXPECT().PostForm(testURL, form).Return(response, nil)

	instagram := service.NewInstagram(httpClientMock)
	result, err := instagram.UsernameIsAvailable(testUsername)
	assert.NoError(t, err)
	assert.True(t, result)
}
