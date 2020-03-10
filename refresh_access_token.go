package jdvop

import (
	"encoding/json"
	"net/url"
)

type RefreshTokenRequest struct {
	client *Client
}

func NewRefreshTokenRequest(client *Client) *RefreshTokenRequest {
	return &RefreshTokenRequest{
		client: client,
	}
}

func (this *RefreshTokenRequest) Method() string {
	return "oauth2/refreshToken"
}

func (this *RefreshTokenRequest) Values() url.Values {
	return this.client.RefreshTokenRequestValues()
}

func (this *RefreshTokenRequest) Refresh() (*Token, error) {
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
