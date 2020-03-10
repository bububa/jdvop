package jdvop

import (
	"encoding/json"
	"net/url"
)

type AccessTokenRequest struct {
	client *Client
}

func NewAccessTokenRequest(client *Client) *AccessTokenRequest {
	return &AccessTokenRequest{
		client: client,
	}
}

func (this *AccessTokenRequest) Method() string {
	return "oauth2/accessToken"
}

func (this *AccessTokenRequest) Values() url.Values {
	return this.client.AccessTokenRequestValues()
}

func (this *AccessTokenRequest) Get() (*Token, error) {
	resp, err := this.client.Do(this)
	if err != nil {
		return nil, err
	}
	var ret Token
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
