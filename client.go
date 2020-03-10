package jdvop

import (
	"crypto/md5"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/bububa/jdvop/internal/debug"
)

type Client struct {
	http.Client
	appId    string
	secret   string
	username string
	passwd   string
	token    IAccessToken
}

func NewClient(appId string, secret string, username string, passwd string) *Client {
	h := md5.New()
	io.WriteString(h, passwd)
	md5Passwd := hex.EncodeToString(h.Sum(nil))
	return &Client{
		appId:    appId,
		secret:   secret,
		username: username,
		passwd:   md5Passwd,
		token:    &DefaultAccessToken{},
	}
}

func (this *Client) SetAccessToken(token IAccessToken) {
	this.token = token
}

func (this *Client) Do(r Request) (json.RawMessage, error) {
	if this.token.Expired() && r.Method() != "oauth2/accessToken" {
		accessTokenReq := NewAccessTokenRequest(this)
		t, err := accessTokenReq.Get()
		if err != nil {
			return nil, err
		}
		this.token.Set(t)
	} else if this.token.RefreshTokenExpired() && r.Method() != "oauth2/accessToken" && r.Method() != "oauth2/refreshToken" {
		refreshTokenReq := NewRefreshTokenRequest(this)
		t, err := refreshTokenReq.Refresh()
		if err != nil {
			return nil, err
		}
		this.token.Set(t)
	}
	values := r.Values()
	if r.Method() != "oauth2/accessToken" && r.Method() != "oauth2/refreshToken" {
		values.Set("token", this.token.Token())
	}
	requestUri := fmt.Sprintf("%s/%s", GATEWAY, r.Method())
	debug.DebugPrintPostMapRequest(requestUri, values)
	resp, err := this.Post(requestUri, "application/x-www-form-urlencoded", strings.NewReader(values.Encode()))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	var ret Response
	err = debug.DecodeJSONHttpResponse(resp.Body, &ret)
	if err != nil {
		return nil, err
	}
	if ret.IsError() {
		return nil, &ret
	}
	return ret.Result, nil
}

func (this *Client) AccessTokenRequestValues() url.Values {
	values := url.Values{}
	ts := time.Now().Format("2006-01-02 15:04:05")
	rawSign := fmt.Sprintf("%s%s%s%s%saccess_token%s", this.secret, ts, this.appId, this.username, this.passwd, this.secret)
	h := md5.New()
	io.WriteString(h, rawSign)
	sign := strings.ToUpper(hex.EncodeToString(h.Sum(nil)))
	values.Set("grant_type", "access_token")
	values.Set("client_id", this.appId)
	values.Set("timestamp", ts)
	values.Set("username", this.username)
	values.Set("password", this.passwd)
	values.Set("sign", sign)
	return values
}

func (this *Client) RefreshTokenRequestValues() url.Values {
	values := url.Values{}
	token := this.token.Get()
	values.Set("refresh_token", token.RefreshToken)
	values.Set("client_id", this.appId)
	values.Set("client_secret", this.secret)
	return values
}
