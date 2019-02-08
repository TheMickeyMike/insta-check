package service

import "net/url"

type RegisterFormResponse struct {
	AccountCreated      bool                `json:"account_created"`
	Errors              *RegisterFormErrors `json:"errors"`
	DryrunPassed        bool                `json:"dryrun_passed"`
	UsernameSuggestions []string            `json:"username_suggestions"`
	Status              string              `json:"status"`
	ErrorType           string              `json:"error_type"`
}

type RegisterFormErrors struct {
	Email    []ValidationError `json:"email"`
	Username []ValidationError `json:"username"`
	Password []ValidationError `json:"password"`
}

type ValidationError struct {
	Message string `json:"message"`
	Code    string `json:"code"`
}

func (form *RegisterFormResponse) UsernameIsOK() bool {
	return len(form.Errors.Username) == 0
}

func NewRegistrationForm(username string) url.Values {
	return url.Values{
		"password":               {"1234"},
		"username":               {username},
		"first_name":             {"first_name"},
		"client_id":              {"client_id"},
		"seamless_login_enabled": {"1"},
		"opt_into_one_tap":       {"false"},
	}
}
