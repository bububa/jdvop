package invoice

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type WaybillRequest struct {
	MarkId string `json:"markId"` // 第三方申请单号：申请发票的唯一id标识
}

func (this *WaybillRequest) Method() string {
	return "api/invoice/waybill"
}

func (this *WaybillRequest) Values() url.Values {
	values := url.Values{}
	values.Set("markId", this.MarkId)
	return values
}

func Waybill(clt *jdvop.Client, markId string) ([]BizInvoiceDelivery, error) {
	resp, err := clt.Do(&WaybillRequest{MarkId: markId})
	if err != nil {
		return nil, err
	}
	var ret []BizInvoiceDelivery
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
