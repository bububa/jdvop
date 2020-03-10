package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetWareReturnJdCompRequest struct {
	OrderId uint64 `json:"jdOrderId"`
	SkuId   uint64 `json:"skuId"`
}

func (this *GetWareReturnJdCompRequest) Method() string {
	return "api/afterSale/getWareReturnJdComp"
}

func (this *GetWareReturnJdCompRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func GetWareReturnJdComp(clt *jdvop.Client, orderId uint64, skuId uint64) ([]ComponentExport, error) {
	resp, err := clt.Do(&GetWareReturnJdCompRequest{OrderId: orderId, SkuId: skuId})
	if err != nil {
		return nil, err
	}
	var ret []ComponentExport
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
