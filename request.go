package tindergo

import (
	"encoding/json"
	"errors"

	"github.com/parnurzeal/gorequest"
)

type HttpRequest interface {
	Get(url string) HttpRequest
	Post(url string) HttpRequest
	Set(name, value string) HttpRequest
	Send(value string) HttpRequest
	End() (string, []error)
}

type Request struct {
	request *gorequest.SuperAgent
}

func NewRequest() *Request {
	return &Request{request: gorequest.New()}
}

func (r *Request) Get(url string) HttpRequest {
	r.request.Get(url)
	return r
}

func (r *Request) Post(url string) HttpRequest {
	r.request.Post(url)
	return r
}

func (r *Request) Set(name, value string) HttpRequest {
	r.request.Set(name, value)
	return r
}

func (r *Request) Send(value string) HttpRequest {
	r.request.Send(value)
	return r
}

func (r *Request) End() (string, []error) {
	_, body, errs := r.request.End()
	return body, errs
}

type TRequest struct {
	requester HttpRequest
	token     string
}

func NewTRequest(token string) *TRequest {
	return &TRequest{
		requester: NewRequest(),
		token:     token,
	}
}

func (t *TRequest) Get(url string) (string, []error) {
	return t.requester.Get(url).
		Set("Host", "api.gotinder.com").
		Set("Authorization", "Token token=\""+t.token+"\"").
		Set("User-Agent", "Tinder/6.9.1 (iPhone; iOS 10.2; Scale/2.00)").
		Set("X-Auth-Token", t.token).
		End()
}

func (t *TRequest) Post(url, data string) (string, []error) {
	return t.requester.Post(url).
		Send(data).
		Set("Host", "api.gotinder.com").
		Set("Authorization", "Token token=\""+t.token+"\"").
		Set("User-Agent", "Tinder/6.9.1 (iPhone; iOS 10.2; Scale/2.00)").
		Set("X-Auth-Token", t.token).
		End()
}

type TAuthRequest struct {
	Token string `json:"token"`
}

type TAuthResponseData struct {
	APIToken  string `json:"api_token"`
	IsNewUser bool   `json:"is_new_user"`
}

type TAuthResponse struct {
	Data TAuthResponseData `json:"data"`
	Meta struct {
		Status int `json:"status"`
	} `json:"meta"`
}

type TAuth struct {
	request TAuthRequest
}

func NewTAuth(token string) *TAuth {

	r := TAuthRequest{
		Token: token,
	}

	return &TAuth{
		request: r,
	}
}

func (a *TAuth) Send() (TAuthResponseData, error) {
	data := TAuthResponseData{}
	r := NewTAuthRequester(NewRequest())
	t, err := r.Authenticate(a.request)
	if err != nil {
		return data, err
	}
	if t.Meta.Status != 200 {
		return data, errors.New("Error while Authenticating")
	}
	return t.Data, nil
}

type TAuthRequester struct {
	requester HttpRequest
}

func NewTAuthRequester(requester HttpRequest) *TAuthRequester {
	return &TAuthRequester{requester: requester}
}

func (tar *TAuthRequester) Authenticate(auth TAuthRequest) (TAuthResponse, error) {
	r := TAuthResponse{}
	url := "https://api.gotinder.com/v2/auth"
	a, err := json.Marshal(auth)
	if err != nil {
		return r, err
	}

	b, errs := tar.requester.Post(url).Send(string(a)).End()
	if errs != nil {
		return r, errs[0]
	}

	err = json.Unmarshal([]byte(b), &r)
	if err != nil {
		return r, err
	}

	return r, nil
}
