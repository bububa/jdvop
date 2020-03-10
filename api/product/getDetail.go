package product

import (
	"encoding/json"
	"net/url"

	"github.com/bububa/jdvop"
)

type GetDetailRequest struct {
	SkuId     string `json:"skuId"`
	QueryExts string `json:"queryExts"`
}

func (this *GetDetailRequest) Method() string {
	return "api/product/getDetail"
}

func (this *GetDetailRequest) Values() url.Values {
	values := url.Values{}
	values.Set("sku", this.SkuId)
	if this.QueryExts != "" {
		values.Set("queryExts", this.QueryExts)
	}
	return values
}

func GetDetail(clt *jdvop.Client, skuId string, queryExts string) (*Sku, error) {
	resp, err := clt.Do(&GetDetailRequest{SkuId: skuId, QueryExts: queryExts})
	if err != nil {
		return nil, err
	}
	var ret Sku
	err = json.Unmarshal(resp, &ret)
	if err != nil {
		return nil, err
	}
	return &ret, nil
}
