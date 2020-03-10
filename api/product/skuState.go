package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type SkuStateRequest struct {
	Sku string `json:"sku"` // 商品编号，支持批量，以“,”（半角）分隔  (最高支持100个商品)
}

func (this *SkuStateRequest) Method() string {
	return "api/product/skuState"
}

func (this *SkuStateRequest) Values() url.Values {
	values := url.Values{}
	values.Set("sku", this.Sku)
	return values
}

func GetSkuState(clt *jdvop.Client, sku string) ([]SkuState, error) {
	resp, err := clt.Do(&SkuStateRequest{Sku: sku})
	if err != nil {
		return nil, err
	}
	var ret []SkuState
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return ret, nil
}
