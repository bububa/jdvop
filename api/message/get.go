package message

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetRequest struct {
	Type string `json:"type"` // 推送类型。支持多个组合，英文逗号间隔。例如1,2,3。支持的参考枚举值：2 商品价格变更 4 商品上下架变更消息 ......
}

func (this *GetRequest) Method() string {
	return "api/message/get"
}

func (this *GetRequest) Values() url.Values {
	values := url.Values{}
	values.Set("type", this.Type)
	return values
}

func Get(clt *jdvop.Client, msgType string) ([]Message, error) {
	resp, err := clt.Do(&GetRequest{Type: msgType})
	if err != nil {
		return nil, err
	}
	var ret []Message
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
