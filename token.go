package jdvop

import (
	"sync/atomic"
	"time"
)

type Token struct {
	AccessToken         string `json:"access_token"`
	RefreshToken        string `json:"refresh_token"`
	Time                int64  `json:"time"`
	ExpiresIn           int64  `json:"expires_in"`
	RefreshTokenExpires int64  `json:"refresh_token_expires"`
}

type IAccessToken interface {
	Token() string
	Get() *Token
	Set(*Token)
	Expired() bool
	RefreshTokenExpired() bool
}

type DefaultAccessToken struct {
	atomic.Value
}

func (this *DefaultAccessToken) Get() *Token {
	t := this.Load()
	if t == nil {
		return nil
	}
	return t.(*Token)
}

func (this *DefaultAccessToken) Set(token *Token) {
	this.Store(token)
}

func (this *DefaultAccessToken) Token() string {
	token := this.Get()
	if token == nil {
		return ""
	}
	return token.AccessToken
}

func (this *DefaultAccessToken) Expired() bool {
	token := this.Get()
	if token == nil {
		return true
	}
	expireTime := time.Unix((token.Time+token.ExpiresIn*1000)/1000, 0)
	return expireTime.Before(time.Now())
}

func (this *DefaultAccessToken) RefreshTokenExpired() bool {
	token := this.Get()
	if token == nil {
		return true
	}
	expireTime := time.Unix(token.RefreshTokenExpires/1000, 0)
	return expireTime.Before(time.Now())
}
