package invoice

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type CancelRequest struct {
	MarkId string `json:"markId"` // 第三方申请单号：申请发票的唯一id标识
}

func (this *CancelRequest) Method() string {
	return "api/invoice/cancel"
}

func (this *CancelRequest) Values() url.Values {
	values := url.Values{}
	values.Set("markId", this.MarkId)
	return values
}

func Cancel(clt *jdvop.Client, markId string) (bool, error) {
	resp, err := clt.Do(&CancelRequest{MarkId: markId})
	if err != nil {
		return false, err
	}
	var ret bool
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return false, err
	}
	return ret, nil
}
