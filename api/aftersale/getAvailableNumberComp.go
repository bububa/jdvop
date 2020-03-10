package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetAvailableNumberCompRequest struct {
	OrderId uint64 `json:"jdOrderId"`
	SkuId   uint64 `json:"skuId"`
}

func (this *GetAvailableNumberCompRequest) Method() string {
	return "api/afterSale/getAvailableNumberComp"
}

func (this *GetAvailableNumberCompRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func GetAvailableNumberComp(clt *jdvop.Client, orderId uint64, skuId uint64) (int, error) {
	resp, err := clt.Do(&GetAvailableNumberCompRequest{OrderId: orderId, SkuId: skuId})
	if err != nil {
		return 0, err
	}
	var ret int
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return 0, err
	}
	return ret, nil
}
