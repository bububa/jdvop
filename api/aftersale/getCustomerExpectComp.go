package aftersale

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetCustomerExpectCompRequest struct {
	OrderId uint64 `json:"jdOrderId"`
	SkuId   uint64 `json:"skuId"`
}

func (this *GetCustomerExpectCompRequest) Method() string {
	return "api/afterSale/getCustomerExpectComp"
}

func (this *GetCustomerExpectCompRequest) Values() url.Values {
	values := url.Values{}
	js, _ := json.Marshal(this)
	values.Set("param", string(js))
	return values
}

func GetCustomerExpectComp(clt *jdvop.Client, orderId uint64, skuId uint64) ([]ComponentExport, error) {
	resp, err := clt.Do(&GetCustomerExpectCompRequest{OrderId: orderId, SkuId: skuId})
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
