package jdvop

import (
	"encoding/json"
	"fmt"
)

type Response struct {
	Code    string          `json:"resultCode,omitempty"`
	Message string          `json:"resultMessage,omitempty"`
	Success bool            `json:"success,omitempty"`
	Result  json.RawMessage `json:"result,omitempty"`
}

func (this *Response) IsError() bool {
	return !this.Success
}

func (this *Response) Error() string {
	return fmt.Sprintf("CODE:%s, MSG:%s", this.Code, this.Message)
}
