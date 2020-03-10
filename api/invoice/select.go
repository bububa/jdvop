package invoice

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type SelectRequest struct {
	MarkId string `json:"markId"` // 第三方申请单号：申请发票的唯一id标识
}

func (this *SelectRequest) Method() string {
	return "api/invoice/select"
}

func (this *SelectRequest) Values() url.Values {
	values := url.Values{}
	values.Set("markId", this.MarkId)
	return values
}

func Select(clt *jdvop.Client, markId string) ([]BizInvoiceDetail, error) {
	resp, err := clt.Do(&SelectRequest{MarkId: markId})
	if err != nil {
		return nil, err
	}
	var ret []BizInvoiceDetail
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
